package setting

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

/**
* @program: src
* @description:
* @author: 占翔昊
* @create 2020-10-05 18:31
**/

type Setting struct {
	vp *viper.Viper
}

/**
初始化配置文件
 */
func NewSettings() (*Setting,error) {
	vp := viper.New()
	path,err := os.Getwd()
	logrus.Info(path)
	if err != nil {
		return nil,err
	}
	vp.SetConfigName("config")
	vp.AddConfigPath(path+"/Pro/gin_standard/configs")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig();err!=nil {
		return nil ,err
	}
	return &Setting{vp},nil
}