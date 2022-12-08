package service

import (
	"cron/internal/lib/define"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局配置项
type (
	Config struct {
		Port    int                        `json:"port"`
		Mod     string                     `json:"mod"`
		Version float64                    `json:"version"`
		Network string                     `json:"network"`
		Mysql   map[define.ChainName]mysql `json:"mysql"`
		Redis   map[string]redis           `json:"redis"`
	}

	// 子项配置
	mysql struct {
		Type       string `json:"type" mapstructure:"type"`
		UserName   string `json:"user_name" mapstructure:"user_name"`
		Password   string `json:"password" mapstructure:"password"`
		Database   string `json:"database" mapstructure:"database"`
		Port       int    `json:"port" mapstructure:"port"`
		Address    string `json:"address" mapstructure:"address"`
		Parameters string `json:"parameters" mapstructure:"parameters"`
		MaxIdle    int    `json:"max_idle" mapstructure:"max_idle"`
		MaxOpen    int    `json:"max_open" mapstructure:"max_open"`
		SSLMode    bool   `json:"ssl_mode" mapstructure:"ssl_mode"`
		Debug      bool   `json:"debug" mapstructure:"debug"`
	}

	redis struct {
		Host     string `json:"host" mapstructure:"host"`
		Port     string `json:"port" mapstructure:"port"`
		Password string `json:"password" mapstructure:"password"`
		DB       int    `json:"database" mapstructure:"database"`
	}
)

var (
	conf            *Config
	defaultConfName = "sys"
)

func GetConfig() *Config {
	return conf
}

//---------------------------内部私有方法---------------------------//

// 导入配置文件
func initConf(confDir, confName string) {
	viper.SetConfigName(confName)
	viper.SetConfigType(define.JsonSuffix)
	viper.AddConfigPath(confDir)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file config: %s", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
}
