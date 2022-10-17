package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	APIPort     string
	DB_ADDRESS  string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	APIKey      string
	TokenSecret string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	viper.Unmarshal(cfg)

	// baca env
	//cfg.APIPort = SetEnv("APIPort", ":6969")
	//cfg.APIKey = SetEnv("APIKey", "kuda-lumping")
	//cfg.TokenSecret = "AbCd3F9H1"

	Cfg = cfg
}

//func SetEnv(key, def string) string {
//	val, ok := os.LookupEnv(key)
//	if !ok {
//		return def
//	}
//
//	return val
//}
