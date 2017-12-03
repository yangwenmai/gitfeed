package main

import (
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
		cli.StringFlag{
			Name:  "include,i",
			Value: "",
			Usage: "Include words. Wildcard pattern matching with support for '?' and '*'",
		},
		cli.StringFlag{
			Name:  "exclude,e",
			Value: "",
			Usage: "Exclude words. Wildcard pattern matching with support for '?' and '*'",
		},
	}
	app.Action = func(c *cli.Context) error {
		cfgInfo := LoadGitFeedCfg(c)
		name := cfgInfo.Username
		maxPage := cfgInfo.MaxPage
		debug := cfgInfo.Debug
		if len(name) > 0 && maxPage > 0 {
			startTime := time.Now()
			// 这里是整个功能的入口
			ReceivedEvents(name, maxPage, debug, c.String("include"), c.String("exclude"))
			cost("Total", startTime)
		}
		return nil
	}
	app.Run(os.Args)
}
