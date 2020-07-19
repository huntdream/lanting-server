package db

import (
	"log"

	"github.com/huntdream/lanting-server/app"
	"github.com/jinzhu/gorm"

	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Initialize database
func Initialize() {
	db, err := gorm.Open("mysql", "root:MySQL888@/lanting?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("Error occurred: ", err)
	}

	app.DB = db
}
