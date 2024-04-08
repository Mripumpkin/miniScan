package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/send", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"a": ctx.Query("a"),
			"b": ctx.Query("b"),
		})
	})
	r.Run("0.0.0.0:8888") // listen and serve on 0.0.0.0:8080
}