package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	BotToken string `mapstructure:"bot_token"`

	OwnerID string `mapstructure:"owner_id"`
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, fmt.Errorf("set up viper: %w", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	expandEnv(reflect.ValueOf(&cfg).Elem())

	return &cfg, nil
}

func expandEnv(value reflect.Value) {
	switch value.Kind() {
	case reflect.String:
		value.SetString(os.ExpandEnv(value.String()))
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			expandEnv(value.Field(i))
		}
	}
}

func setUpViper() error {
	viper.AutomaticEnv()

	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
