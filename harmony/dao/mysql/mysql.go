package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"harmony/common/parseyaml"
	"harmony/model"
)

var DB *gorm.DB

func InitMysql() {
	username := parseyaml.Vp.GetString("db.username")
	password := parseyaml.Vp.GetString("db.password")
	addr := parseyaml.Vp.GetString("db.addr")
	port := parseyaml.Vp.GetInt("db.port")
	dbname := parseyaml.Vp.GetString("db.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	err = DB.AutoMigrate(&model.Channel{}, &model.Server{}, &model.Message{}, &model.Member{}, model.Conversation{}, &model.DirectMessage{}, &model.Profile{})
	if err != nil {
		panic(err)
		return
	}
}
