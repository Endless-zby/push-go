package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Type  string
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
	}
	Sqlite struct {
		File string
	}
}

type ApnsConfig struct {
	AuthKeyFile string
	TeamID      string
	KeyID       string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Apns     ApnsConfig
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
