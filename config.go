package utils

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Instances map[string]Instance
	Producers map[string]Producer
	Consumers map[string]Consumer
}
type Instance struct {
	ID      string
	Listen  string
	Produce string
}
type Producer struct {
	ID      string
	Listen  string
	Produce string
	Topic   string
	Address string
}
type Consumer struct {
	ID      string
	Produce string
	Topic   string
	Address string
}

func (c Config) Read(path string) error {
	if _, err := toml.DecodeFile(path, &c); err != nil {
		return err
	}
	return nil
}
