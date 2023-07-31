package internal

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func InitConfigs() error {
	if err := initLogger(); err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func ReadConfigFromYAML(filename string, config interface{}) error {
	// Read the YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal the YAML data into the LogConfig struct
	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}

	return nil
}
