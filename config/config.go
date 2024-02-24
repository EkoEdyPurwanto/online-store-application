package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type (
	DbConfig struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		Driver   string
	}

	ApiConfig struct {
		ApiHost string
		ApiPort string
	}

	FileConfig struct {
		FilePath string
	}

	TokenConfig struct {
		ApplicationName  string
		JwtSignatureKey  []byte
		JwtSigningMethod jwt.SigningMethodHMAC
		ExpirationToken  int
	}
)

type Config struct {
	DbConfig
	ApiConfig
	FileConfig
	TokenConfig
}

// Constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	return err
}
