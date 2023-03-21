package main

type Config struct {
	Source      DatabaseConfig   `mapstructure:"source"`
	Destination []DatabaseConfig `mapstructure:"destination"`
}

type DatabaseConfig struct {
	ID       string `mapstructure:"id" json:"id"`
	IP       string `mapstructure:"ip"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Name     string `mapstructure:"name"`
}
