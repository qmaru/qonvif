package onvif

import (
	"fmt"

	"qonvif/configs"
	"qonvif/configs/models"

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
	var deviceInfo models.DeviceInfo

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

func (basic *OnvifBasic) getDeviceInformation(dev *goonvif.Device, method any) (*DeviceInfo, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("tds:GetDeviceInformationResponse")
	return &DeviceInfo{
		Manufacturer:    xmlData.SelectElement("tds:Manufacturer").Text(),
		Model:           xmlData.SelectElement("tds:Model").Text(),
		FirmwareVersion: xmlData.SelectElement("tds:FirmwareVersion").Text(),
		SerialNumber:    xmlData.SelectElement("tds:SerialNumber").Text(),
		HardwareId:      xmlData.SelectElement("tds:HardwareId").Text(),
	}, nil
}

func (basic *OnvifBasic) getNetworkInterfaces(dev *goonvif.Device, method any) (*DeviceNetwork, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("tds:GetNetworkInterfacesResponse")
	xmlInterface := xmlData.SelectElement("tds:NetworkInterfaces")
	xmlInterfaceInfo := xmlInterface.SelectElement("tt:Info")

	return &DeviceNetwork{
		Iface:   xmlInterfaceInfo.SelectElement("Name").Text(),
		MacAddr: xmlInterfaceInfo.SelectElement("HwAddress").Text(),
	}, nil
}

func (basic *OnvifBasic) getProfiles(dev *goonvif.Device, method any) ([]*ProfileData, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	profileData := make([]*ProfileData, 0)

	xmlData := body.SelectElement("trt:GetProfilesResponse")
	xmlProfiles := xmlData.SelectElements("trt:Profiles")
	for _, xmlProfile := range xmlProfiles {
		profile := &ProfileData{
			Name:  xmlProfile.SelectElement("tt:Name").Text(),
			Token: xmlProfile.SelectAttr("token").Value,
		}
		profileData = append(profileData, profile)
	}
	return profileData, nil
}

func (basic *OnvifBasic) getStreamUri(dev *goonvif.Device, method any) (*StreamData, error) {
	body, err := handleMethod(dev, method)
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("trt:GetStreamUriResponse")
	xmlMediaUri := xmlData.SelectElement("trt:MediaUri")

	return &StreamData{
		Url: xmlMediaUri.SelectElement("tt:Uri").Text(),
	}, nil
}

func (basic *OnvifBasic) GetDeviceData() (*DeviceData, error) {
	deviceInfo, err := basic.getDeviceInformation(basic.device, device.GetDeviceInformation{})
	if err != nil {
		return nil, err
	}

	deviceNetwork, err := basic.getNetworkInterfaces(basic.device, device.GetNetworkInterfaces{})
	if err != nil {
		return nil, err
	}
	deviceNetwork.Addr = basic.xAddr

	deviceData := &DeviceData{
		Info:    *deviceInfo,
		Network: *deviceNetwork,
	}

	return deviceData, nil
}

func (basic *OnvifBasic) GetProfiles() ([]*ProfileData, error) {
	return basic.getProfiles(basic.device, media.GetProfiles{})
}

func (basic *OnvifBasic) GetStreamUri(token string) (*StreamData, error) {
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

func (basic *OnvifBasic) usePTZAbsoluteMove(dev *goonvif.Device, x, y float64) error {
	_, err := handleMethod(dev, ptz.AbsoluteMove{
		Position: onvif.PTZVector{
			PanTilt: onvif.Vector2D{
				X:     x,
				Y:     y,
				Space: xsd.AnyURI("DS"),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (basic *OnvifBasic) usePTZRelativeMove(dev *goonvif.Device, x, y float64) error {
	_, err := handleMethod(dev, ptz.RelativeMove{
		Translation: onvif.PTZVector{
			PanTilt: onvif.Vector2D{
				X:     x,
				Y:     y,
				Space: xsd.AnyURI("DS"),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (basic *OnvifBasic) getPTZStatus() (*PtzStatusData, error) {
	body, err := handleMethod(basic.device, ptz.GetStatus{})
	if err != nil {
		return nil, err
	}

	xmlData := body.SelectElement("tptz:GetStatusResponse")
	xmlStatus := xmlData.SelectElement("tptz:PTZStatus")
	xmlPosition := xmlStatus.SelectElement("tt:Position")
	xmlPanTilt := xmlPosition.SelectElement("tt:PanTilt")
	return &PtzStatusData{
		X: xmlPanTilt.SelectAttr("x").Value,
		Y: xmlPanTilt.SelectAttr("y").Value,
	}, nil
}

func (basic *OnvifBasic) PTZStatus() (*PtzStatusData, error) {
	ptzStatus, err := basic.getPTZStatus()
	if err != nil {
		return nil, err
	}
	return ptzStatus, nil
}

func (basic *OnvifBasic) PTZGoToAnyAbsolute(x, y float64) (*PtzStatusData, error) {
	err := basic.usePTZAbsoluteMove(basic.device, x, y)
	if err != nil {
		return nil, err
	}

	ptzStatus, err := basic.getPTZStatus()
	if err != nil {
		return nil, err
	}
	return ptzStatus, nil
}

func (basic *OnvifBasic) PTZGoToAnyRelative(x, y float64) (*PtzStatusData, error) {
	err := basic.usePTZRelativeMove(basic.device, x, y)
	if err != nil {
		return nil, err
	}

	ptzStatus, err := basic.getPTZStatus()
	if err != nil {
		return nil, err
	}
	return ptzStatus, nil
}
