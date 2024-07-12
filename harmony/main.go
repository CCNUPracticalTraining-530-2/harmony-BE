package main

import (
	"harmony/common/parseyaml"
	"harmony/dao/mysql"
	"harmony/router"
)

func main() {
	parseyaml.GetYaml()
	mysql.InitMysql()
	r := router.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
