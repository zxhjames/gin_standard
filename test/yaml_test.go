package test

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"testing"
)

/**
* @program: 测试go读取yaml文件
* @description:
* @author: 占翔昊
* @create 2020-10-05 19:19
**/

func Test_Load_yaml(t *testing.T) {
	config := viper.New()
	p,_ := os.Getwd()
	config.AddConfigPath(p)     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}else {
		//logrus.Info(config)
	}
	//打印文件读取出来的内容:
	logrus.Info(config.Get("Database.Username"))
}