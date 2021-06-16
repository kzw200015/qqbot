package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var BotConfig struct {
	NickName      []string `yaml:"nick_name"`
	CommandPrefix string   `yaml:"command_prefix"`
	SuperUsers    []string `yaml:"super_users"`
	Url           string   `yaml:"url"`
	AccessToken   string   `yaml:"access_token"`
}

func init() {
	bytes, err := os.ReadFile("./config.yml")
	if err != nil {
		logrus.Fatalln(err)
	}

	err = yaml.Unmarshal(bytes, &BotConfig)
	if err != nil {
		logrus.Fatalln(err)
	}
}
