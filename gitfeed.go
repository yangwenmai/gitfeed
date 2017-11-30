package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gitfeed"
	app.Usage = "Check GitHub Newsfeed."
	app.Version = "0.0.1"
	app.Author = "maiyang"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Value: "",
			Usage: "Load configuration from `FILE` (default:~/.gitfeed/gitfeed.ini)",
		},
		cli.StringFlag{
			Name:  "user,u",
			Value: "",
			Usage: "Github username",
		},
	}
	app.Action = func(c *cli.Context) error {
		cfgInfo := LoadGitFeedCfg(c)
		name := cfgInfo.Username
		maxPage := cfgInfo.MaxPage
		if len(name) > 0 && maxPage > 0 {
			// 这里是整个功能的入口
			startTime := time.Now()
			ReceivedEvents(name, maxPage)
			fmt.Printf("Total cost( %v )", time.Now().Sub(startTime))
		}
		return nil
	}
	app.Run(os.Args)
}
