package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	data := "aaa"

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
		// ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
