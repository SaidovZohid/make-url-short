package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort  string
	RedisAddr string
	MongoDB   MongoDB
}

type MongoDB struct {
	Host           string
	Port           string
	User           string
	Password       string
	Database       string
	UserCollection string
	UrlCollection  string
}

func Load(path string) (*Config, error) {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, err
	}
	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		HttpPort:  conf.GetString("HTTP_PORT"),
		RedisAddr: conf.GetString("REDIS_ADDR"),
		MongoDB: MongoDB{
			Host:           conf.GetString("MONGODB_HOST"),
			Port:           conf.GetString("MONGODB_PORT"),
			User:           conf.GetString("MONGODB_USER"),
			Password:       conf.GetString("MONGODB_PASSWORD"),
			Database:       conf.GetString("MONGODB_DATABASE"),
			UserCollection: conf.GetString("MONGODB_COLLECTION_USER"),
			UrlCollection:  conf.GetString("MONGODB_COLLECTION_URL"),
		},
	}
	return &cfg, nil
}
