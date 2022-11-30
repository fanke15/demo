package test

import (
	"cron/internal/lib/define"
	"cron/internal/lib/service"
	"flag"
)

var configName = flag.String("f", "sys_local", "the config file")

func init() {
	service.New(define.One).SetConfName(*configName).
		SetLogDirName("logs").SetLogMsgPrefix("service", "logTest").Init()
}
