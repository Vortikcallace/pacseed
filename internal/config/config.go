package config

import (
	"errors"
	"fmt"
	"log/slog"

	"pacseed/internal/config/models"

	"github.com/spf13/viper"
)

var UConfig *models.Config

var ErrRepoNotFound = errors.New("repo not found")

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

func GetRepo(name string) (models.Repo, error) {
	if UConfig == nil {
		return models.Repo{}, errors.New("config not loaded")
	}
	for _, r := range UConfig.Repos {
		if r.Name == name {
			return r, nil
		}
	}
	return models.Repo{}, ErrRepoNotFound
}

func GetAURRepos() []models.Repo {
	if UConfig == nil {
		return nil
	}
	out := make([]models.Repo, 0, len(UConfig.Repos))
	for _, r := range UConfig.Repos {
		if r.AUR {
			out = append(out, r)
		}
	}
	return out
}

func GetRepos() []models.Repo {
	slog.Debug("get repos from config")
	if UConfig == nil {
		return nil
	}
	out := make([]models.Repo, len(UConfig.Repos))
	copy(out, UConfig.Repos)
	return out
}

func PrintRepo(name string) {
	slog.Debug("get repo by name")
	r, err := GetRepo(name)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	source := "official"
	if r.AUR {
		source = "AUR"
	}

	fmt.Printf("name:    %s\n", r.Name)
	fmt.Printf("source:  %s\n", source)
	fmt.Printf("version: %s\n", r.Version)
}

func PrintRepos() {
	repos := GetRepos()
	if len(repos) == 0 {
		fmt.Println("no repos")
		return
	}

	for i, r := range repos {
		source := "official"
		if r.AUR {
			source = "AUR"
		}

		fmt.Printf("[%d] %s (%s) v%s\n", i+1, r.Name, source, r.Version)
	}
}

func PrintAURRepos() {
	repos := GetAURRepos()
	if len(repos) == 0 {
		fmt.Println("no AUR repos")
		return
	}

	for i, r := range repos {
		fmt.Printf("[%d] %s v%s\n", i+1, r.Name, r.Version)
	}
}
