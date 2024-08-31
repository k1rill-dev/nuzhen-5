package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
)

type Config struct {
	Env              string `yaml:"env"`
	DatabaseHost     string `yaml:"databaseHost"`
	DatabasePort     int    `yaml:"databasePort"`
	DatabaseUser     string `yaml:"databaseUser"`
	DatabaseName     string `yaml:"databaseName"`
	DatabasePassword string `yaml:"databasePassword"`
	DatabaseSSLMode  string `yaml:"databaseSSLMode"`
}

func fetchConfigPath(filename string) string {
	path, err := filepath.Abs(filename)
	if err != nil {
		errString := fmt.Sprintf("Error getting absolute path of config file %s: %s\n", filename, err)
		panic(errString)
	}
	return path
}

func MustLoadConfig(filename string) *Config {
	configPath := fetchConfigPath(filename)
	if configPath == "" {
		panic("config path not set")
	}
	fmt.Println(configPath)
	if q, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println(q)
		panic("config path does not exist")
	}
	config := &Config{}
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		panic(err)
	}
	return config
}
