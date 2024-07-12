package parseyaml

import "github.com/spf13/viper"

var Vp *viper.Viper

func GetYaml() {
	Vp = viper.New()
	Vp.AddConfigPath("./config") // 路径
	Vp.SetConfigName("config")   // 名称
	Vp.SetConfigType("yaml")     // 类型
	err := Vp.ReadInConfig()     // 读取配置
	if err != nil {
		panic(err)
		return
	}
}
