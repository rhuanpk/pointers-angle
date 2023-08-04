package models

// API is the struct that store all API configuration.
type API struct {
	Port *uint16 `mapstructure:"port"`
}

// DB is the struct that store all DB configuration.
type DB struct {
	Host           *string `mapstructure:"host"`
	User           *string `mapstructure:"user"`
	Pass           *string `mapstructure:"pass"`
	Name           *string `mapstructure:"name"`
	Port           *uint16 `mapstructure:"port"`
	SSL            *string `mapstructure:"ssl"`
	TZ             *string `mapstructure:"tz"`
	MaxIdleConns   *uint8  `mapstructure:"max_idle_conns"`
	MaxOpenConns   *uint8  `mapstructure:"max_open_conns"`
	MaxLifeSecConn *uint8  `mapstructure:"max_life_sec_conn"`
}
