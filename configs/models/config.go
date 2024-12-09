package models

type DeviceInfo struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
}

type DeviceConfig struct {
	Name   string     `toml:"name" json:"name"`
	Device DeviceInfo `toml:"device" json:"device"`
}

type ServerConfig struct {
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	ApiKey string `toml:"api_key"`
}

type Config struct {
	Desc    string         `toml:"desc"`
	Debug   bool           `toml:"debug"`
	Server  ServerConfig   `toml:"server"`
	Devices []DeviceConfig `toml:"devices"`
}
