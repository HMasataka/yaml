package main

import (
	"io/ioutil"

	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
)

type Config []Component

type Component struct {
	Func string      `json:"func" yaml:"func"`
	Args interface{} `json:"args" yaml:"args"`
}

func LoadYaml(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func main() {
	config, err := LoadYaml("config.yaml")
	if err != nil {
		panic(err)
	}
	pp.Printf("%v", config)
}
