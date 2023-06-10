package main

import (
	"github.com/Nojus0/go-gin-sql/Mysql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

func Entrypoint(releaseMode bool) *gin.Engine {
	if err := Mysql.Connect(); err != nil {
		log.Fatalln("Failed to Connect to database:", err)
	}

	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"time":   time.Now().Local(),
		})
	})

	return r
}
