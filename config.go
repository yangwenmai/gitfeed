package main

import (
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

// LoadGitFeedCfg 加载 gitfeed 配置文件
func LoadGitFeedCfg(c *cli.Context) GitFeedCfg {
	username := ""
	maxPage := 1
	debug := false

	// 从指令中加载参数
	if len(c.String("user")) > 0 {
		username = c.String("user")
	}
	if c.Int("max_page") > 0 {
		maxPage = c.Int("max_page")
	}
	if c.Bool("debug") {
		debug = true
	}

	if len(username) == 0 {
		basePath, _ := homedir.Dir()
		path := basePath + "/.gitfeed/gitfeed.ini"
		if len(c.String("config")) > 0 {
			// 读取到你设置的配置文件
			path = c.String("config")
		}
		// 加载配置文件
		cfg, err := goconfig.LoadConfigFile(path)
		if err != nil {
			log.Fatalf("无法加载配置文件，请创建并配置参数：%s", err)
		}
		username, err = cfg.GetValue("GitHub Newsfeed", "user")
		if err != nil {
			log.Fatalf("无法获取键值（%s）：%s", "username", err)
		}
		maxPage, err = cfg.Int("GitHub Newsfeed", "max_page")
		if err != nil {
			log.Fatalf("无法获取键值（%s）：%s", "max_page", err)
		}
		debug, err = cfg.Bool("GitHub Newsfeed", "debug")
		if err != nil {
			log.Fatalf("无法获取键值（%s）：%s", "debug", err)
		}
	}

	cfgInfo := GitFeedCfg{}
	cfgInfo.Username = username
	cfgInfo.MaxPage = maxPage
	cfgInfo.Debug = debug
	return cfgInfo
}

// GitFeedCfg gitfeed config
type GitFeedCfg struct {
	Username string `json:"username"`
	MaxPage  int    `json:"max_page"`
	Debug    bool   `json:"debug"`
}
