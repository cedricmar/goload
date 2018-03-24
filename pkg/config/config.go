package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config type wraps a *json.RawMessage
type Config struct {
	MainDir string
	Color   string
}

// LoadConfig gets the config
func LoadConfig() *Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var c *Config
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// GetMainDir a value from the Config
func (c *Config) GetMainDir() string {
	return c.MainDir
}

// GetColor get a value from the Config
func (c *Config) GetColor() string {
	return c.Color
}
