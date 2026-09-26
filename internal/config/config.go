package config

import (
	"log/slog"

	"pacseed/internal/config/models"

	"github.com/spf13/viper"
)

var UConfig *models.Config

func LoadConfig() {
	viper.SetConfigName("my_template")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./testdata/mytestdata")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Error reading config file", "error", err)
		return
	}

	var ucfg models.Config
	if err := viper.Unmarshal(&ucfg); err != nil {
		slog.Error("Error parsing config", "error", err)
		return
	}

	UConfig = &ucfg
	slog.Debug("Config loaded successfully",
		"path", viper.ConfigFileUsed(),
		"repos", len(ucfg.Repos),
	)
}

func saveConfig() error {
	if err := viper.WriteConfig(); err != nil {
		slog.Error("Error writing config file", "error", err)
		return err
	}
	slog.Debug("Config saved successfully")
	return nil
}
