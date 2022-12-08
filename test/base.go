package test

import (
	"cron/internal/lib/define"
	"cron/internal/lib/service"
	"flag"
)

var configName = flag.String("f", "sys_local", "the config file")

func init() {
	sn := service.New(define.One)
	sn.SetConfName(*configName)
	sn.SetLogDirName("logs")
	sn.SetLogMsgPrefix("service", "logTest")
	sn.Init()

}
