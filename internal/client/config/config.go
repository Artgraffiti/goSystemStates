// CLIENT Config
package config

import "github.com/spf13/viper"

type Config struct {
	// ServerAddr - адрес сервера
	ServerAddr string `mapstructure:"SERVER_ADDR"`
}

func LoadConfig(path string) (config Config, err error) {
	// Config default values
	viper.SetDefault("SERVER_ADDR", "127.0.0.1:8080")

	viper.AddConfigPath(path)
	viper.SetConfigName("main")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
