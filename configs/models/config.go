package models

type DeviceInfo struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type DeviceConfig struct {
	Name   string     `toml:"name"`
	Device DeviceInfo `toml:"device"`
}

type ServerConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type Config struct {
	Desc    string         `toml:"desc"`
	Debug   bool           `toml:"debug"`
	Server  ServerConfig   `toml:"server"`
	Devices []DeviceConfig `toml:"devices"`
}
