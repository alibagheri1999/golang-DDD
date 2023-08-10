package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

const (
	// ProductionMode indicates mode is release.
	ProductionMode = "production"
	// DebuggingMode indicates mode is debug.
	DebuggingMode = "debugging"
)

var cfg *Config

func Init() {
	cfg = new(Config)
	log.Println()
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath("../config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("error loading default configs: ", err)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.MergeInConfig(); err != nil {
		log.Println("no config file found. Using defaults and environment variables.")
	}
	if err := v.UnmarshalExact(&cfg); err != nil {
		log.Fatalf("invalid config schema: %v", err)
	}
	//if err := validator.New().Struct(cfg); err != nil {
	//	log.Fatal("invalid config: %v", err)
	//}
	// generate a unique id for this instance
	//cfg.App.InstanceID =
}

func Set(newConfig *Config) {
	cfg = newConfig
}

func Get() *Config {
	return cfg
}

func (c *Config) IsProduction() bool {
	return c.App.Env == ProductionMode
}

func (c *Config) IsDebugging() bool {
	return c.App.Env == DebuggingMode
}
