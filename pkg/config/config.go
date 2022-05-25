package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Web struct {
		AppPort int `mapstructure:"app_port"`
	}

	Proxy struct {
		AppAddress string `mapstructure:"app_address"`
		WafPort    int    `mapstructure:"waf_port"`
	}

	DB struct {
		User    string
		Pass    string
		Address string
		Port    int
	}

	Modsec struct {
		Conf    string `mapstructure:"coroza_conf"`
		Setup   string `mapstructure:"modsec_ruleset_setup"`
		Ruleset string `mapstructure:"modesec_ruleset"`
	}
}

func Read() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/goalkeeper/")
	viper.AddConfigPath("$HOME/.goalkeeper")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return c
}
