package utils

import (
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

func (c *Config) Read(path string) error {
	if _, err := toml.DecodeFile(path, &c); err != nil {
		return err
	}
	return nil
}
