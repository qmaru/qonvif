package onvif

type DeviceInfo struct {
	Manufacturer    string `json:"manufacturer"`
	Model           string `json:"model"`
	FirmwareVersion string `json:"firmware_version"`
	SerialNumber    string `json:"serial_number"`
	HardwareId      string `json:"hardware_id"`
}

type DeviceNetwork struct {
	Iface   string `json:"iface"`
	Addr    string `json:"addr"`
	MacAddr string `json:"mac_addr"`
}

type DeviceData struct {
	Info    DeviceInfo    `json:"info"`
	Network DeviceNetwork `json:"network"`
}
