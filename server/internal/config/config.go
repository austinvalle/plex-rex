package config

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
)

// Config defines the shape of the whole api configuration
type Config struct {
	TheMovieDb TheMovieDbConfig `toml:"themoviedb"`
}

// TheMovieDbConfig defines the shape of themoviedb configuration
type TheMovieDbConfig struct {
	BaseURL string `toml:"base_url"`
	APIKey  string `toml:"api_key"`
}

// GetConfig retrieves config information from config.toml file
func GetConfig() Config {
	_, workspacePath, _, _ := runtime.Caller(1)
	var configFile = path.Join(path.Dir(workspacePath), "./config.toml")
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing: ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
