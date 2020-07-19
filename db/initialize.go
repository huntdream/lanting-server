package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
)

//Ci structure
type Ci struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
}

//Init database
func Init() {
	var count int
	data := []Ci{}

	ci, err := ioutil.ReadFile("./db/ci.json")

	if err != nil {
		log.Fatalln(err)
	}

	app.DB.Table("articles").Count(&count)

	json.Unmarshal([]byte(ci), &data)

	for _, item := range data {
		fmt.Println(item)
		app.DB.Table("articles").Create(&model.Article{
			Title:   item.Title,
			Content: strings.Join(item.Paragraphs, ""),
		})
	}

}
