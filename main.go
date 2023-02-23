package main

import (
	"awesomeProject1/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// An example about DSN. It should read from environment variables
	dsn := "test:test@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	utils.InitDBConnection(dsn, gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	r := gin.Default()
	r.POST("/phone", phoneCreateHandler)
	r.Run()
}
