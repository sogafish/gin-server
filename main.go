package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DataBaseInit
func dbInit() {
	db, err := gorm.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic("gormOpen Error")
	}

	db.AutoMigrate(&Todo{})

	defer db.Close()
}

// Insert
func dbCreate(text string, status string) {
	db, err := gorm.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic("gormOpen Error dbCreate")
	}

	db.Create(&Todo{Text: text, Status: status})

	defer db.Close()
}

func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic("gormOpen Error dbUpdate")
	}
	var todo Todo

	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)

	db.Close()
}

func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic("gormOpen Error dbDelete")
	}
	var todo Todo

	db.First(&todo, id)
	db.Delete(&todo)

	db.Close()
}

// get All todos Records
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic("gormOpen Error dbGetAll")
	}

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()

	return todos
}

// get All todos Record
func dbGet(id int) Todo {
	db, err := gorm.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic("gormOpen Error dbGet")
	}

	var todo Todo
	db.First(&todo, id)
	db.Close()

	return todo
}

func main() {
	router := gin.Default()

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>_start")

	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()

		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.POST("/add", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")

		dbCreate(text, status)
		ctx.Redirect(302, "/")
	})

	router.GET("detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")

		id, err := strconv.Atoi(n)

		if err != nil {
			panic(err)
		}

		todo := dbGet(id)
		ctx.HTML(200, "todo.html", gin.H{"todo": todo})
	})

	router.POST("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")

		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		dbUpdate(id, text, status)

		todo := dbGet(id)
		ctx.HTML(200, "todo.html", gin.H{"todo": todo})
	})

	router.GET("/delete_c/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")

		id, err := strconv.Atoi(n)

		if err != nil {
			panic(err)
		}

		todo := dbGet(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)

		if err != nil {
			panic(err)
		}

		dbDelete(id)

		ctx.Redirect(302, "/")
	})

	router.Run()
}
