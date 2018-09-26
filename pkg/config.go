package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const Version = "v0.2.0"
var config = &Config{}
type Config struct {
	ServerInfo ServerInfo `yaml:"serverInfo"`
}

type ServerInfo struct {
	EnvName string `yaml:"envName"`
}

func ReadConfig(configFile string) {
	buf, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		panic(err)
	}
	config = &c
}
