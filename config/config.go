package config

type DatabaseConfig struct {
	Schema       string `mapstructure:"schema"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"database_name"`
	SSLMode      string `mapstructure:"sslmode"`
}
