package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/lovoo/opsgenie-cardiogram/cardiogram"

	"gopkg.in/yaml.v2"
)

type config struct {
	URL      string        `yaml:"url"`
	APIKey   string        `yaml:"api_key"`
	Timeout  time.Duration `yaml:"timeout"`
	Interval time.Duration `yaml:"interval"`
	Targets  []target      `yaml:"targets"`
}

type target struct {
	Name       string `yaml:"name"`
	URL        string `yaml:"url"`
	StatusCode int    `yaml:"status_code"`
}

func (c *config) readConfig(path string) *config {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read config file")
	}

	err = yaml.Unmarshal(f, c)
	if err != nil {
		log.Fatal("Cannot parse config file")
	}

	return c
}

func main() {
	var c config
	configPath := flag.String("config", "config.yml", "path to the configuration file")
	flag.Parse()

	c.readConfig(*configPath)

	h := cardiogram.Heartbeat{
		Client: &http.Client{
			Timeout: c.Timeout,
		},
		URL:    c.URL,
		APIKey: c.APIKey,
	}

	log.Println("Starting Opsgenie Cardiogram")

	for range time.Tick(c.Interval) {
		for _, t := range c.Targets {
			go h.Check(t.URL, t.StatusCode, t.Name)
		}
	}
}
