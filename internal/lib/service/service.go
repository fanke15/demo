// Package service 应用依赖服务入口封装
package service

import (
	"cron/internal/lib/define"
	"cron/pkg"
	"fmt"
	"path/filepath"
	"strings"
)

type Service struct{}

var (
	defaultServiceDirDepth = define.Two
	defaultServiceDir      = define.StrNull
)

// New 设置目录
func New(depth ...int) *Service {
	var (
		dir, _       = filepath.Abs(define.StrDat)
		paths        = strings.Split(dir, define.Forwardslash)
		serviceDepth = define.Zero
	)
	if len(depth) > define.Zero {
		serviceDepth = depth[define.Zero]
	}
	if len(paths) > serviceDepth {
		defaultServiceDirDepth = serviceDepth
	}
	defaultServiceDir = pkg.AnySliceToStr(define.Forwardslash, paths[:len(paths)-defaultServiceDirDepth]...)
	return &Service{}
}
func (s *Service) SetConfName(confName string) {
	if confName != define.StrNull {
		defaultConfName = confName
	}
	defaultConfName = pkg.AnySliceToStr(define.Forwardslash, "conf", defaultConfName)
	fmt.Println("当前使用配置文件路径:", pkg.AnySliceToStr(define.Forwardslash, defaultServiceDir, defaultConfName))
	return
}
func (s *Service) SetLogDirName(name string) {
	defaultLogDirName = name
	return
}
func (s *Service) SetLogMsgPrefix(prefixKVs ...interface{}) {
	defaultKV = prefixKVs
	return
}
func (s *Service) SetDBNames(names ...define.ChainName) {
	dbNames = append(dbNames, names...)
}

// Init 初始化各项服务
func (s *Service) Init() {
	initConf(defaultServiceDir, defaultConfName)
	initLog()
	for _, v := range dbNames {
		initDB(v)
	}
	for _, v := range redisNames {
		initRedis(v)
	}
}

//---------------------------内部私有方法---------------------------//
