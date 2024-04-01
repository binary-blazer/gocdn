package config

import (
	"gocdn/types"
)

func GetConfig() types.Config {
	return types.Config{
		Webserver: types.Webserver{
			Port: 8080,
		},
		Database: types.Database{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "password",
			Database: "gocdn",
		},
		UploadTypes: []string{
			"image/jpeg",
			"image/png",
			"image/gif",
		},
		BigTypes: []string{
			"logo",
			"banner",
		},
		MaxUploadSize: 1024 * 1024 * 5,
	}
}
