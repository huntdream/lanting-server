package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/huntdream/lanting-server/app"
)

//Initialize database
func Initialize() {
	name := app.Config.Database.Name
	username := app.Config.Database.User
	passwd := app.Config.Database.Passwd

	dsl := username + ":" + passwd + "@/" + name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsl)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	app.DB = db
}
