package main

import (
	"cron/internal/lib/define"
	"cron/internal/lib/service"
	"flag"
)

var configDepth = flag.Int("depth", define.Two, "the config file depth")
var configName = flag.String("f", "sys", "the config file")

func main() {
	flag.Parse()

	init := service.New(*configDepth)
	init.SetConfName(*configName)
	init.SetLogMsgPrefix("network", define.EthereumName)
	init.SetLogDirName("logfiles")
	init.Init()
}
