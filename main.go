package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// func dbInit() {
// 	db, err := gorm.Open("sqlite3", "test.sqlite3")

// 	if err != nil {
// 		panic("gormOpen Error")
// 	}

// 	db.Create(&Todo{Text: text, Status: status})

// 	defer db.Close()
// }

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	data := "aaa"

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}
