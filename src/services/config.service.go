package services

import (
	"gorestapi/src/entities"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func ConfigInit(config *entities.Configuration) error {
	var environment string
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "prod" {
		environment = "development"
	} else {
		environment = "production"
	}

	viper.SetConfigName("config." + environment)
	viper.AddConfigPath("./src/config")
	viper.AutomaticEnv()

	// Change _ underscore in env to . dot notation in viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Read config
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}
