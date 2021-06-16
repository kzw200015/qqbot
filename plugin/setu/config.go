package setu

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var config struct {
	ApiAddr string `yaml:"api_addr"`
	Proxy   string `yaml:"proxy"`
	R18     string `yaml:"r18"`
}

func init() {
	bytes, err := os.ReadFile("./setu.yml")
	if err != nil {
		logrus.Fatalln(err)
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		logrus.Fatalln(err)
	}
}
