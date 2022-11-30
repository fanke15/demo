package service

import (
	"cron/internal/lib/define"
	"cron/pkg"
	"fmt"
	unlog "log"
	"os"
	"sync"
	"time"

	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dbs           sync.Map
	defaultDBName define.ChainName = "default"
	dbNames                        = []define.ChainName{defaultDBName}
)

func GetMysql(dbName ...define.ChainName) *gorm.DB {
	var name = defaultDBName
	if len(dbName) > define.Zero {
		name = dbName[define.Zero]
	}
	if dbTemp, ok := dbs.Load(name); ok {
		return dbTemp.(*gorm.DB)
	}
	return nil
}

//---------------------------内部私有方法---------------------------//

// 初始化连接
func initDB(dbName ...define.ChainName) {
	var (
		name = defaultDBName
		err  error
		db   *gorm.DB
	)
	if len(dbName) > define.Zero {
		name = dbName[define.Zero]
	}
	oldDB, _ := dbs.Load(name)

	// 新建连接
	if err = pkg.Retry(func() error {
		db, err = openOrm(name)
		if err != nil {
			return err
		}
		return nil
	}, define.Three, define.Three*time.Second); err != nil {
		panic(err)
	}

	dbs.Delete(name)
	dbs.Store(name, db)
	if oldDB != nil {
		if dbTemp, _ := oldDB.(*gorm.DB).DB(); dbTemp != nil {
			dbTemp.Close()
		}
	}
	GetLog().Info("mysql conn ok !", "dbName", name)
}

func openOrm(dbname define.ChainName) (*gorm.DB, error) {
	var (
		info = GetConfig().Mysqls[dbname] // 获取配置信息进行连接
		lc   = logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logger.Silent,          // Log level
			IgnoreRecordNotFoundError: true,
			Colorful:                  false, // 禁用彩色打印，日志平台会打印出颜色码，影响日志观察
		}
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?%s&parseTime=true",
			info.UserName,
			info.Password,
			info.Address,
			info.Port,
			info.Database,
			info.Parameters,
		)
	)
	if info.Debug {
		lc.LogLevel = logger.Info
	}

	orm, err := gorm.Open(driver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			unlog.New(os.Stdout, "[GORM] >> ", 64|unlog.Ldate|unlog.Ltime),
			lc,
		),
	})
	if err != nil {
		return nil, err
	}
	db, err := orm.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(info.MaxIdle) // 连接池的空闲数大小
	db.SetMaxOpenConns(info.MaxOpen) // 最大打开连接数
	db.SetConnMaxLifetime(define.NegativeOne)
	return orm, nil
}
