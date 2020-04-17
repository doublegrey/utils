package utils

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Instances map[string]Instance
	Producers map[string]Producer
	Consumers map[string]Consumer
	Params    map[string]string
}
type Instance struct {
	ID      string
	Listen  string
	Produce string
	Params  map[string]string
}
type Producer struct {
	ID      string
	Listen  string
	Produce string
	Topic   string
	Address string
	Params  map[string]string
}
type Consumer struct {
	ID      string
	Produce string
	Topic   string
	Address string
	Params  map[string]string
}

func (c *Config) Read(path string) error {
	if _, err := toml.DecodeFile(path, &c); err != nil {
		return err
	}
	return nil
}
