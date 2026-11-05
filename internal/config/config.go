package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"pacseed/internal/config/models"

	"github.com/spf13/viper"
)

var UConfig *models.Config
var ConfigPath string
var ConfigName string

var ErrRepoNotFound = errors.New("repo not found")

func LoadConfig(cfgname string, cfgpath string) error {
	ConfigName = cfgname
	ConfigPath = cfgpath

	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var ucfg models.Config
	if err := viper.Unmarshal(&ucfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	UConfig = &ucfg
	slog.Debug("Config loaded successfully", "path", viper.ConfigFileUsed(), "repos", len(ucfg.Repos))
	return nil
}

func SaveConfig() error {
	if UConfig == nil {
		return errors.New("nothing to save")
	}

	if ConfigPath == "" {
		ConfigPath, _ = os.UserHomeDir()
	}
	if ConfigName == "" {
		ConfigName = "repos"
	}

	fullPath := filepath.Join(ConfigPath, ConfigName+".yaml")

	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("File %s already exists. Overwrite? [y/N]: ", fullPath)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			return errors.New("aborted by user")
		}
	}

	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigPath)

	if err := viper.MergeConfigMap(map[string]interface{}{"repos": UConfig.Repos}); err != nil {
		return fmt.Errorf("failed to merge config: %w", err)
	}

	if err := viper.WriteConfigAs(fullPath); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	slog.Debug("Config saved successfully", "path", fullPath)
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

		fmt.Printf("[%d] %s (%s)\n", i+1, r.Name, source)
	}
}

func PrintAURRepos() {
	repos := GetAURRepos()
	if len(repos) == 0 {
		fmt.Println("no AUR repos")
		return
	}

	for i, r := range repos {
		fmt.Printf("[%d] %s\n", i+1, r.Name)
	}
}
