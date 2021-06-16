package setu

import (
	"flag"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/shell"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
)

func init() {
	zero.OnCommand("setu").Handle(func(ctx *zero.Ctx) {
		var num uint
		var keyword string
		var isSingle bool
		err := parseArgs(&num, &keyword, &isSingle, ctx.State["args"].(string))
		if err != nil {
			logrus.Error(err)
			ctx.Send(message.Text(err.Error()))
			return
		}

		pics, err := getPics(num, keyword)
		if err != nil {
			ctx.Send(message.Text(err.Error()))
			return
		}

		if isSingle { //单张发送
			for _, pic := range pics {
				ctx.Send(message.Image(pic))
			}
		} else { //合并发送
			var messageSegments []message.MessageSegment
			for _, pic := range pics {
				messageSegments = append(messageSegments, message.Image(pic))
			}
			ctx.Send(messageSegments)
		}
	})
}

func getPics(num uint, keyword string) ([]string, error) {
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"proxy":   config.Proxy,
		"r18":     config.R18,
		"num":     strconv.Itoa(int(num)),
		"keyword": keyword,
	}).Get(config.ApiAddr)
	if err != nil {
		return nil, err
	}

	r := gjson.Get(resp.String(), "data.#.urls.original")

	var pics []string
	for _, i := range r.Array() {
		pics = append(pics, i.String())
	}

	return pics, nil
}

func parseArgs(num *uint, keyword *string, isSingle *bool, args string) error {
	fs := flag.FlagSet{}
	fs.UintVar(num, "n", 1, "")
	fs.StringVar(keyword, "k", "", "")
	fs.BoolVar(isSingle, "s", false, "")
	arguments := shell.Parse(args)

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	return nil
}
