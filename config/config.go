package config

import (
	"io/ioutil"
	"log"
	"os"
	yaml "gopkg.in/yaml.v2"
)

// Config for the app
type Config struct {
	KafkaTopic     string `yaml:"kafka.topic"`
	KafkaVerifySsl bool   `yaml:"kafka.verify"`
	KafkaBrokers   string `yaml:"kafka.brokers"`
	KafkaVerbose   bool   `yaml:"kafka.verbose"`
	KafkaCaFile    string `yaml:"kafka.ca"`
	KafkaKeyFile   string `yaml:"kafka.key"`
	KafkaCertFile  string `yaml:"kafka.certificate"`
	ListenAddr     string `yaml:"web.listen.addr"`
	TelemetryPath  string `yaml:"web.telemetry.path"`
}

// GetConfig returns the app's configuration described on config.yaml on root
func GetConfig() *Config {
	cfg := &Config{}

	_ ,err := os.Stat("./config.yaml")
	if err == nil{
		yamlFile, err := ioutil.ReadFile("./config.yaml")

		if err != nil {
			log.Printf("Road config.yaml failed: #%v ", err)
		}

		err = yaml.Unmarshal(yamlFile, cfg)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}else{
		log.Println("Config File config.yaml not found.")
	}

	return cfg
}
