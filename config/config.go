package config

import "os"

type Config struct {
	APIPort     string
	APIKey      string
	TokenSecret string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	// baca env
	cfg.APIPort = os.Getenv("APIPort")
	cfg.APIKey = os.Getenv("APIKey")
	cfg.TokenSecret = "AbCd3F9H1"

	Cfg = cfg
}
