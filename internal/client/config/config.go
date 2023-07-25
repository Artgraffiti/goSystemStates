// CLIENT Config
package config

import "github.com/spf13/viper"

type Config struct {
	// ServerAddr - адрес сервера
	ServerAddr string `mapstructure:"SERVER_ADDR"`

	// GRPCServerAddr - адрес GRPC сервера (по умолчанию: 127.0.0.1:50051)
	GRPCServerAddr string `mapstructure:"GRPC_SERVER_ADDR"`

	// UUID - универсальный уникальный индификатор пользователя
	UUID string `mapstructure:"UUID"`
}

func LoadConfig(path string) (config Config, err error) {
	// Config default values
	viper.SetDefault("SERVER_ADDR", "127.0.0.1:8080")
	viper.SetDefault("GRPC_SERVER_ADDR", "127.0.0.1:50051")
	viper.SetDefault("UUID", "1412fe28-06b0-408b-832d-f2d5aa4792a5")

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
