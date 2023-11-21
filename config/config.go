package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type FileConfig struct{}

type TokenConfig struct{}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.ApiConfig.ApiPort == "" || c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Name == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" {
		fmt.Println(c.ApiConfig.ApiPort, c.DbConfig.Driver, c.DbConfig.Host, c.DbConfig.Name, c.DbConfig.Port, c.DbConfig.User)
		return errors.New("all environment variables required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := conf.readConfig(); err != nil {
		return nil, err
	}
	return conf, nil
}
