package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/structJSON_test", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "foo"
		msg.Message = "これは構造体をJSONで返すためのテストです。"
		msg.Number = 1111
		context.JSON(http.StatusOK, msg)
	})

	_ = r.Run(":9000")
}
