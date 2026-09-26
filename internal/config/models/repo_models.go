package models

type Repo struct {
	Name    string `mapstructure:"name"`
	AUR     bool   `mapstructure:"aur"`
	Version string `mapstructure:"version"`
}

type Config struct {
	Repos []Repo `mapstructure:"repos"`
}
