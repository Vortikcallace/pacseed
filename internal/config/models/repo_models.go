package models

type Repo struct {
	Name string `mapstructure:"name"`
	AUR  bool   `mapstructure:"aur"`
}

type Config struct {
	Repos []Repo `mapstructure:"repos"`
}
