package main

import (
	"github.com/kzw200015/qqbot/config"
	_ "github.com/kzw200015/qqbot/log"
	_ "github.com/kzw200015/qqbot/plugin/music"
	_ "github.com/kzw200015/qqbot/plugin/repeater"
	_ "github.com/kzw200015/qqbot/plugin/setu"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

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
