package config

import "github.com/spf13/viper"

type AppConfig struct {
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Brokers  string
	Topic    string
	GroupId  string
	ClientId string
}

func GetConfig() *AppConfig {
	var appConfig AppConfig

	v := viper.New()
	v.SetConfigFile("resources/config.yml")

	if err := v.ReadInConfig(); err != nil {
		panic("error while reading configuration")
	}

	err := v.Unmarshal(&appConfig)
	if err != nil {
		panic("error while marshalling configuration")
	}

	return &appConfig
}
