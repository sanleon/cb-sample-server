package main

import (
	"log"
	"testing"
)

func TestReadConfig(t *testing.T) {
	ReadConfig("../config.yaml")

	if config.ServerInfo.EnvName != "local" {
		log.Fatal("EnvName is not local")
	}
}
