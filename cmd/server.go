package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func serve() {

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"version": version,

		})
	})
	log.Fatal(r.Run(":8080"))
}
