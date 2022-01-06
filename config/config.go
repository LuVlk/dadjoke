package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	basePath = "basePath"
)

func Init() error {
	viper.SetDefault(basePath, "https://icanhazdadjoke.com/")

	viper.AutomaticEnv()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	viper.AddConfigPath(path.Join(home, ".dadjoke"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return fmt.Errorf("failed to initialize config - %w", err)
	}

	return nil
}

func BasePath() string {
	return viper.Get(basePath).(string)
}
