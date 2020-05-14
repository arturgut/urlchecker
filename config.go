package main

import (
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/go-yaml/yaml.v2"
)

// Config struct
type Config struct {
	Server struct {
		Port int
	}

	Client struct {
		SkipSSL bool
		Timeout time.Duration
		Period  time.Duration
	}

	Urls []string
}

var config Config

func loadConfiguration(filename string) {
	log.Info("INFO: Loading configuration...")
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("An error occured: ", err)
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Error("An error occured: ", err)
		panic(err)
	}
	log.Info("Config has been sucessfully loaded.")
	log.Debug("Config map: ", config)
}
