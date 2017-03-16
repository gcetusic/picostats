package main

import (
	"flag"
	"fmt"
	"os"
)

type CliParser struct {
	Command string
}

func (c *CliParser) parseCli() {
	help := flag.Bool("h", false, "")
	helpLong := flag.Bool("help", false, "")

	listen := flag.String("l", "", "")
	listenLong := flag.String("listen", "", "")

	redis := flag.String("r", "", "")
	redisLong := flag.String("redis", "", "")

	dev := flag.Bool("d", false, "")
	devLong := flag.Bool("dev", false, "")

	flag.Parse()

	if *help || *helpLong {
		c.printHelpMessage()
		os.Exit(0)
	}

	if len(*listen) > 0 {
		conf.ListenAddr = *listen
	} else if len(*listenLong) > 0 {
		conf.ListenAddr = *listenLong
	}

	if len(*redis) > 0 {
		conf.RedisUrl = *redis
	} else if len(*redisLong) > 0 {
		conf.RedisUrl = *redisLong
	}

	if *dev || *devLong {
		conf.Dev = true
	}

	for _, arg := range os.Args {
		if arg == "server" || arg == "worker" || arg == "serverww" {
			c.Command = arg
		}
	}
}

func (c *CliParser) printHelpMessage() {
	helpMessage := `PicoStats ` + VERSION + ` - (C) 2017 Tihomir Piha

Usage: picostats [COMMAND] [FLAGS]

Commands:

    server     - run in server mode together with worker (default)
    worker     - run worker without server
    serverww   - run in server mode without worker

Flags:

    -l,   --listen=ADDRESS      server listen address (HOST:PORT)
    -r,   --redis=REDIS_URL     Redis connection url
    -d,   --dev                 run in development mode (more logging)
    -h,   --help                print this message
`
	fmt.Println(helpMessage)
}

func initCli() {
	clip = &CliParser{Command: "server"}
	clip.parseCli()
}
