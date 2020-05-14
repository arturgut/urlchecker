package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/go-yaml/yaml.v2"
)

// Config struct
type Config struct {
	Server struct {
		Port int
	}

	Client struct {
		SkipSSL bool
		Timeout int
	}

	Urls []string
}

var config Config

func loadConfiguration(filename string) {

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("DEBUG: Config:", config.Server.Port)
	fmt.Println("DEBUG: Config:", config.Client.SkipSSL)
	fmt.Println("DEBUG: Config:", config.Client.Timeout)
	for _, value := range config.Urls {
		fmt.Println("DEBUG: Config:", value)
	}
}
