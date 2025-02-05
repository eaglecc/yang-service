package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	ServerPort int     `mapstructure:"serverPort" json:"serverPort" yaml:"serverPort"`
	Db         MysqlDB `mapstructure:"database" json:"database" yaml:"database"`
}

type MysqlDB struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DBName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	PassWord string `mapstructure:"password" json:"password" yaml:"password"`
	Myconfig string `mapstructure:"myconfig" json:"myconfig" yaml:"myconfig"`
}

// 配置文件中解析命令行参数
func InitConfig() {
	Conf = new(Config)
	v := viper.New()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	v.AddConfigPath(path + "/config")
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	// 监听config文件变动
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		err2 := v.Unmarshal(&Conf)
		if err2 != nil {
			panic(fmt.Errorf("Unmarshal change config data，err：%v\n", err2))
		}
	})
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错，err：%v\n", err))
	}
	// 解析
	err = v.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("Unmarshal config data，err：%v\n", err))
	}
}
