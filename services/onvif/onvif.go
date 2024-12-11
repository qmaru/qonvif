package onvif

import (
	"fmt"

	"qonvif/configs"
	configModels "qonvif/configs/models"

	"qonvif/services/onvif/models"

	goonvif "github.com/use-go/onvif"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd"
	"github.com/use-go/onvif/xsd/onvif"
)

type OnvifBasic struct {
	xAddr  string
	device *goonvif.Device
}

func NewClient(name string) (*OnvifBasic, error) {
	var deviceInfo configModels.DeviceInfo

	for _, d := range configs.Config.Devices {
		if d.Name == name {
			deviceInfo = d.Device
		}
	}

	xAddr := fmt.Sprintf("%s:%d", deviceInfo.Host, deviceInfo.Port)
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{Xaddr: xAddr, Username: deviceInfo.Username, Password: deviceInfo.Password})
	if err != nil {
		return nil, err
	}

	return &OnvifBasic{
		xAddr:  xAddr,
		device: dev,
	}, nil
}

func (basic *OnvifBasic) getDeviceInformation(dev *goonvif.Device, method any) (*models.DeviceInfo, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("tds:GetDeviceInformationResponse")
	return &models.DeviceInfo{
		Manufacturer:    xmlData.SelectElement("tds:Manufacturer").Text(),
		Model:           xmlData.SelectElement("tds:Model").Text(),
		FirmwareVersion: xmlData.SelectElement("tds:FirmwareVersion").Text(),
		SerialNumber:    xmlData.SelectElement("tds:SerialNumber").Text(),
		HardwareId:      xmlData.SelectElement("tds:HardwareId").Text(),
	}, nil
}

func (basic *OnvifBasic) getNetworkInterfaces(dev *goonvif.Device, method any) (*models.DeviceNetwork, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("tds:GetNetworkInterfacesResponse")
	xmlInterface := xmlData.SelectElement("tds:NetworkInterfaces")
	xmlInterfaceInfo := xmlInterface.SelectElement("tt:Info")

	return &models.DeviceNetwork{
		Iface:   xmlInterfaceInfo.SelectElement("Name").Text(),
		MacAddr: xmlInterfaceInfo.SelectElement("HwAddress").Text(),
	}, nil
}

func (basic *OnvifBasic) getProfiles(dev *goonvif.Device, method any) ([]*models.ProfileData, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	profileData := make([]*models.ProfileData, 0)

	xmlData := body.SelectElement("trt:GetProfilesResponse")
	xmlProfiles := xmlData.SelectElements("trt:Profiles")
	for _, xmlProfile := range xmlProfiles {
		profile := &models.ProfileData{
			Name:  xmlProfile.SelectElement("tt:Name").Text(),
			Token: xmlProfile.SelectAttr("token").Value,
		}
		profileData = append(profileData, profile)
	}
	return profileData, nil
}

func (basic *OnvifBasic) getStreamUri(dev *goonvif.Device, method any) (*models.StreamUriData, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("trt:GetStreamUriResponse")
	xmlMediaUri := xmlData.SelectElement("trt:MediaUri")

	return &models.StreamUriData{
		Url: xmlMediaUri.SelectElement("tt:Uri").Text(),
	}, nil
}

func (basic *OnvifBasic) getPTZStatus() (*models.PtzStatusData, error) {
	body, err := handleMethod(basic.device, ptz.GetStatus{})
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("tptz:GetStatusResponse")
	xmlStatus := xmlData.SelectElement("tptz:PTZStatus")
	xmlPosition := xmlStatus.SelectElement("tt:Position")
	xmlPanTilt := xmlPosition.SelectElement("tt:PanTilt")
	xmlZoom := xmlPosition.SelectElement("tt:Zoom")
	return &models.PtzStatusData{
		X: xmlPanTilt.SelectAttr("x").Value,
		Y: xmlPanTilt.SelectAttr("y").Value,
		Z: xmlZoom.SelectAttr("x").Value,
	}, nil
}

func (basic *OnvifBasic) GetDeviceData() (*models.DeviceData, error) {
	deviceInfo, err := basic.getDeviceInformation(basic.device, device.GetDeviceInformation{})
	if err != nil {
		return nil, err
	}

	deviceNetwork, err := basic.getNetworkInterfaces(basic.device, device.GetNetworkInterfaces{})
	if err != nil {
		return nil, err
	}
	deviceNetwork.Addr = basic.xAddr

	deviceData := &models.DeviceData{
		Info:    *deviceInfo,
		Network: *deviceNetwork,
	}

	return deviceData, nil
}

func (basic *OnvifBasic) GetProfiles() ([]*models.ProfileData, error) {
	return basic.getProfiles(basic.device, media.GetProfiles{})
}

func (basic *OnvifBasic) GetStreamUri(token string) (*models.StreamUriData, error) {
	method := media.GetStreamUri{
		StreamSetup: onvif.StreamSetup{
			Stream: "RTP_unicast",
			Transport: onvif.Transport{
				Protocol: "RTSP",
			},
		},
		ProfileToken: onvif.ReferenceToken(token),
	}

	return basic.getStreamUri(basic.device, method)
}

func (basic *OnvifBasic) setPTZPanTilt(x, y float64) onvif.Vector2D {
	return onvif.Vector2D{
		X:     x,
		Y:     y,
		Space: xsd.AnyURI("DS"),
	}
}

func (basic *OnvifBasic) setPTZZoom(z float64) onvif.Vector1D {
	return onvif.Vector1D{
		X:     z,
		Space: xsd.AnyURI("DS"),
	}
}

func (basic *OnvifBasic) PTZStatus() (*models.PtzStatusData, error) {
	return basic.getPTZStatus()
}

// PTZGoToAnyAbsolute
//
//	x: horizontal
//	y: vertical
//	z: Zoom
func (basic *OnvifBasic) PTZGoToAnyAbsolute(x, y, z float64) (*models.PtzStatusData, error) {
	_, err := handleMethod(basic.device, ptz.AbsoluteMove{
		Position: onvif.PTZVector{
			PanTilt: basic.setPTZPanTilt(x, y),
			Zoom:    basic.setPTZZoom(z),
		},
	})

	if err != nil {
		return nil, err
	}

	return basic.getPTZStatus()
}

// PTZGoToAnyRelative
//
//	x: horizontal
//	y: vertical
//	z: Zoom
func (basic *OnvifBasic) PTZGoToAnyRelative(x, y, z float64) (*models.PtzStatusData, error) {
	_, err := handleMethod(basic.device, ptz.RelativeMove{
		Translation: onvif.PTZVector{
			PanTilt: basic.setPTZPanTilt(x, y),
			Zoom:    basic.setPTZZoom(z),
		},
	})

	if err != nil {
		return nil, err
	}

	return basic.getPTZStatus()
}
