package config

var C *Config

type Config struct {
	Path     string
	Range    []int
	Endpoint string
}
