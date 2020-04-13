package utils

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Instances map[string]instance
	Producers map[string]producer
	Consumers map[string]consumer
}
type instance struct {
	ID      string
	Listen  string
	Produce string
}
type producer struct {
	ID      string
	Listen  string
	Produce string
	Topic   string
	Address string
}
type consumer struct {
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
