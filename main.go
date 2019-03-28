package main

import (
	"github.com/ElusiveSpirit/go-store/config"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := InitServer()

	r.Run(":8000")
}

func InitServer() *gin.Engine {
	db := config.InitDB()
	env := &config.Env{DB: db}

	r := gin.Default()

	config.CreateRoutes(r, env)

	return r
}
