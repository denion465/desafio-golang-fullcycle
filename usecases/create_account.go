package usecases

import (
	"log"
	"time"

	"github.com/denion465/desafio-golang-fullcycle/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func CreateAccount(c *gin.Context) {
	account := models.Account{}
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println("account", account.Account_number)
	db := c.MustGet("db").(*gorm.DB)
	accountDB := models.Account{}
	db.Find(&accountDB, account)
	if accountDB.Account_number != "" {
		c.JSON(400, gin.H{
			"error": "Account " + account.Account_number + " already exists",
		})
		return
	}
	account.ID = uuid.NewV4().String()
	account.Balance = 1000
	account.Created_at = time.Now()
	db.Create(&account)
	c.JSON(201, gin.H{
		"ID":             account.ID,
		"account_number": account.Account_number,
		"balance":        account.Balance,
	})
}
