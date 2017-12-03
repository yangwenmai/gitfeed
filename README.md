# GitFeed #
[![Build Status](https://travis-ci.org/yangwenmai/gitfeed.svg?branch=master)](https://travis-ci.org/yangwenmai/gitfeed) [![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/gitfeed)](https://goreportcard.com/report/github.com/yangwenmai/gitfeed)  [![Documentation](https://godoc.org/github.com/yangwenmai/gitfeed?status.svg)](http://godoc.org/github.com/yangwenmai/gitfeed) [![Coverage Status](https://coveralls.io/repos/github/yangwenmai/gitfeed/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/gitfeed?branch=master) [![GitHub issues](https://img.shields.io/github/issues/yangwenmai/gitfeed.svg)](https://github.com/yangwenmai/gitfeed/issues) [![license](https://img.shields.io/github/license/yangwenmai/gitfeed.svg?maxAge=2592000)](https://github.com/yangwenmai/gitfeed/LICENSE) [![Release](https://img.shields.io/github/release/yangwenmai/gitfeed.svg?label=Release)](https://github.com/yangwenmai/gitfeed/releases)

[中文版本](README_ZH.md)

Check GitHub Newsfeed via the command-line in Go, insipred by [GitFeed](https://github.com/ritiek/GitFeed).

Newsfeed includes all the news from people you are following on GitHub, repositories you are watching, etc.

All news you would find on your GitHub dashboard.

Base on [Github Evets](https://developer.github.com/v3/activity/events/#list-public-events-that-a-user-has-received)

>**List events that a user has received**
>These are events that you've received by watching repos and following users. If you are authenticated as the given user, you will see private events. Otherwise, you'll only see public events.

## Screenshots ##

![gitfeed screenshots](docs/gitfeed.png)

## Installation ##

1. `go get github.com/yangwenmai/gitfeed`.

2. install gitfeed by source

    - `cd $GOPATH/src/github.com/yangwenmai/`
    - `git clone https://github.com/yangwenmai/gitfeed.git`
    - `cd gitfeed`
    - `go build && ./gitfeed`

## Usage ##

Run it using gitfeed

The first time you launch gitfeed, it will ask you for GitHub username and set it as the default username to fetch news for.

You can even fetch news for any other user provided you know their GitHub username.

Full list of supported options:

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

You can modify the default configuration by editing ~/.gitfeed/gitfeed.ini

## Roadmap ##

What are you want?

Please create a new Issue for me.

## License ##

License: MIT License