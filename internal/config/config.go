package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var Config *AppConfig

type AppConfig struct {
	HTTP     HTTPConfig `yaml:"http"`
	Log      LogConfig  `yaml:"log"`
	Sentry   Sentry     `yaml:"sentry"`
	Services Services   `yaml:"services"`
}

type HTTPConfig struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type LogConfig struct {
	File          string `yaml:"file"`
	Level         string `yaml:"level"`
	SyslogEnabled bool   `yaml:"syslog_enable"`
}

type Sentry struct {
	Enabled bool   `yaml:"enabled"`
	Key     string `yaml:"key"`
}

type Services struct {
	Artists       string `yaml:"artists"`
	Feed          string `yaml:"feed"`
	Subscriptions string `yaml:"subscriptions"`
	Users         string `yaml:"users"`
}

func InitConfig(filepath string) error {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	if err := Load(data); err != nil {
		return err
	}

	log.Infof("Config loaded from %v.", filepath)
	return nil
}

func Load(data []byte) error {
	cfg := AppConfig{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}
	Config = &cfg
	return nil
}
