package main

import (
	"github.com/denion465/desafio-golang-fullcycle/usecases"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	db := ConnectDatabase()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/bank-accounts", usecases.CreateAccount)
	r.POST("/bank-accounts/transfer", usecases.TransferBetweenAccounts)
	r.Run()
}

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
