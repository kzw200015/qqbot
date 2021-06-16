package main

import (
	"github.com/kzw200015/qqbot/config"
	_ "github.com/kzw200015/qqbot/plugin/setu"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func init() {
	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	zero.Run(zero.Config{
		NickName:      config.BotConfig.NickName,
		CommandPrefix: config.BotConfig.CommandPrefix,
		SuperUsers:    config.BotConfig.SuperUsers,
		Driver: []zero.Driver{
			driver.NewWebSocketClient(config.BotConfig.Url, config.BotConfig.AccessToken),
		},
	})

	select {}
}
