package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func InitConfigs() error {
	if err := initLogger(); err != nil {
		hlog.Error(err.Error())
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
