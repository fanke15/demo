package service

import (
	"context"
	"cron/internal/lib/define"
	"cron/pkg"
	rc "github.com/go-redis/redis/v9"
	"sync"
	"time"
)

var (
	caches           sync.Map
	defaultRedisName = "default"
	redisNames       = []string{defaultRedisName}
)

func GetRedis(redisName ...string) *rc.Client {
	var name = defaultRedisName
	if len(redisName) > define.Zero {
		name = redisName[define.Zero]
	}
	if cacheTemp, ok := caches.Load(name); ok {
		return cacheTemp.(*rc.Client)
	}
	return nil
}

// ---------------------------内部私有方法---------------------------//
// 初始化连接
func initRedis(redisName ...string) {
	var (
		name  = defaultRedisName
		err   error
		cache *rc.Client
	)
	if len(redisName) > define.Zero {
		name = redisName[define.Zero]
	}
	oldRedis, _ := caches.Load(name)

	// 新建连接
	if err = pkg.Retry(func() error {
		cache, err = connRedis(name)
		if err != nil {
			return err
		}
		return nil
	}, define.Three, define.Three*time.Second); err != nil {
		panic(err)
	}
	caches.Delete(name)
	caches.Store(name, cache)

	if oldRedis != nil {
		oldRedis.(*rc.Client).Close()
	}

	GetLog().Info("redis conn ok !", "redisName", name)
}

func connRedis(name string) (*rc.Client, error) {
	var (
		info = GetConfig().Redis[name] // 获取配置信息进行连接
	)
	opt := &rc.Options{
		Addr:     pkg.AnySliceToStr(define.StrColon, info.Host, info.Port),
		DB:       info.DB,
		Password: info.Password,
	}
	conn := rc.NewClient(opt)
	if err := conn.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return conn, nil
}
