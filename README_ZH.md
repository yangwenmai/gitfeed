# GitFeed #
[![Build Status](https://travis-ci.org/yangwenmai/gitfeed.svg?branch=master)](https://travis-ci.org/yangwenmai/gitfeed) [![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/gitfeed)](https://goreportcard.com/report/github.com/yangwenmai/gitfeed)  [![Documentation](https://godoc.org/github.com/yangwenmai/gitfeed?status.svg)](http://godoc.org/github.com/yangwenmai/gitfeed) [![Coverage Status](https://coveralls.io/repos/github/yangwenmai/gitfeed/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/gitfeed?branch=master) [![GitHub issues](https://img.shields.io/github/issues/yangwenmai/gitfeed.svg)](https://github.com/yangwenmai/gitfeed/issues) [![license](https://img.shields.io/github/license/yangwenmai/gitfeed.svg?maxAge=2592000)](https://github.com/yangwenmai/gitfeed/LICENSE) [![Release](https://img.shields.io/github/release/yangwenmai/gitfeed.svg?label=Release)](https://github.com/yangwenmai/gitfeed/releases)

在命令行查看 Github 用户的 Newsfeed，包括 Github 用户 follow 的人，以及 watch 的仓库等所有动态，这些动态你都能够在你的 Github 控制面板上看到。

Newsfeed 是基于 [Github Events API]( https://developer.github.com/v3/activity/events/#list-public-events-that-a-user-has-received)。

初始版本的灵感来源于：[GitFeed-Python](https://github.com/ritiek/GitFeed).

>**列出一个用户接收到的事件**
>你能接收到你 follow 的人，以及 watch 项目的动态事件。如果用户授权了，你还可以看到私有事件，否则你只能看到公开的时间。

## 功能截屏 ##

![gitfeed screenshots](docs/gitfeed.png)

## 安装 ##

1. `go get github.com/yangwenmai/gitfeed`.

2. 源码安装

    - `cd $GOPATH/src/github.com/yangwenmai/`
    - `git clone https://github.com/yangwenmai/gitfeed.git`
    - `cd gitfeed`
    - `go build && ./gitfeed`

## 用法 ##

用 `gitfeed` 来运行。

在运行 `gitfeed` 之前, 你必须在 `~/.gitfeed/gitfeed.ini` 中配置好相应参数`username`,`max_page`,`debug`，当然你也可以指定其他 Github 用户名，以获取他们的公开动态。

命令参数:

```shell
NAME:
   gitfeed - Check GitHub Newsfeed.

USAGE:
   gitfeed [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   maiyang

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE     Load configuration from FILE (default:~/.gitfeed/gitfeed.ini)
   --user value, -u value     Github username
   --include value, -i value  Include words. Wildcard pattern matching with support for '?' and '*'
   --exclude value, -e value  Exclude words. Wildcard pattern matching with support for '?' and '*'
   --help, -h                 show help
   --version, -v              print the version
```

![](docs/wxpay.jpg)

## 下一步？ ##

你想要什么呢？请你提 Issue 给我吧
如果有bug怎么办？也请提 Issue 给我。

## 如何参与/贡献？ ##

1. Fork 此项目
2. 克隆你自己的项目到你本地 （git clone https://github.com/your_github_name/gitfeed.git）
2. 创建你新的 feature 分支 (git checkout -b my_feature)
3. 添加并提交你的修改内容 (git commit -am 'Add some feature')
4. 推送到你自己项目的远端 feature 分支 (git push origin my_feature)
5. 创建一个新的 PR（Pull Request）

## License ##

License: MIT License