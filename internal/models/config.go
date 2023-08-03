package models

// Config is the struct that store all app configuration.
type Config struct {
	API struct {
		Port *uint16 `mapstructure:"port"`
	} `mapstructure:"api"`
	DB struct {
		Name *string `mapstructure:"name"`
		Host *string `mapstructure:"host"`
		Port *uint16 `mapstructure:"port"`
		User *string `mapstructure:"user"`
		Pass *string `mapstructure:"pass"`
	} `mapstructure:"db"`
}
