package music

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	zero.OnCommand("music").Handle(func(ctx *zero.Ctx) {
		args := ctx.State["args"].(string)
		if args == "" {
			ctx.Send(message.Text("需要歌曲名"))
			return
		}
		music, err := getMusic(args)
		if err != nil {
			logrus.Errorln(err)
			return
		}

		ctx.Send(message.Music("163", music))
	})
}

func getMusic(name string) (int64, error) {
	client := resty.New()
	resp, err := client.R().SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.66").
		SetQueryParams(map[string]string{
			"type": "1",
			"s":    name,
		}).Get("https://music.163.com/api/search/get")
	if err != nil {
		return 0, err
	}

	return gjson.ParseBytes(resp.Body()).Get("result.songs.0.id").Int(), nil
}
