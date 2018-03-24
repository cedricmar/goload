package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config type wraps a *json.RawMessage
type Config struct {
	vals map[string]*json.RawMessage
}

/*
// Get the config
file, err := ioutil.ReadFile("config.json")
if err != nil {
	log.Fatal(err)
}

var rawConfig map[string]*json.RawMessage
if err = json.Unmarshal(file, &rawConfig); err != nil {
	log.Fatal(err)
}

config := fmt.Sprintf("./%s", strings.Trim(string(*rawConfig["main_dir"]), "\""))
*/

// LoadConfig gets the config
func LoadConfig() *Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}
	err = json.Unmarshal(file, &c.vals)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}

// Get a value from the Config
func (c *Config) Get(k string) string {
	return fmt.Sprintf("%s", *c.vals[k])
}
