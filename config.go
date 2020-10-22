package utils

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Instances map[string]Instance
	Producers map[string]Producer
	Consumers map[string]Consumer
	Params    map[string]interface{}
}
type Instance struct {
	ID      string
	Listen  string
	Produce string
	Params  map[string]interface{}
}
type Producer struct {
	ID      string
	Listen  string
	Produce string
	Topic   string
	Address string
	Params  map[string]interface{}
}
type Consumer struct {
	ID      string
	Produce string
	Topic   string
	Address string
	Params  map[string]interface{}
}

// Parse flow.toml config
//
// use CONFIG_PATH env var if no parameters given
func (c *Config) Parse(path ...string) error {
	config := os.Getenv("CONFIG_PATH")
	if len(path) > 0 {
		config = path[0]
	}
	if _, err := toml.DecodeFile(config, &c); err != nil {
		return err
	}
	return nil
}
