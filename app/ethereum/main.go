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

	service.SetConfPath(*configName, *configDepth)
	service.NewConfig()
}
