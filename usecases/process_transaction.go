package usecases

import (
	"github.com/denion465/desafio-golang-fullcycle/dtos"
	"github.com/denion465/desafio-golang-fullcycle/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func TransferBetweenAccounts(c *gin.Context) {
	transactionDto := dtos.Transaction{}
	transaction := models.Transaction{}
	toAccount := models.Account{}
	fromAccount := models.Account{}
	if err := c.ShouldBindJSON(&transactionDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.First(&toAccount, "account_number = ?", transactionDto.To)
	db.First(&fromAccount, "account_number = ?", transactionDto.From)
	if toAccount.ID == "" {
		c.JSON(400, gin.H{
			"error": "Account " + transactionDto.To + " does not exist",
		})
		return
	}
	if fromAccount.ID == "" {
		c.JSON(400, gin.H{
			"error": "Account " + transactionDto.From + " does not exist",
		})
		return
	}
	if fromAccount.Balance <= 0 {
		c.JSON(400, gin.H{
			"error": "Account " + transactionDto.From + " has insufficient funds",
		})
		return
	}
	dbUpdate := db.Model(&toAccount).Update("balance", toAccount.Balance+transactionDto.Amount)
	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"error": dbUpdate.Error.Error(),
		})
		return
	}
	if toAccount.Account_number != fromAccount.Account_number {
		dbU := db.Model(&fromAccount).Update("balance", fromAccount.Balance-transactionDto.Amount)
		if dbU.Error != nil {
			c.JSON(400, gin.H{
				"error": dbU.Error.Error(),
			})
			return
		}
	}
	transaction.ID = uuid.NewV4().String()
	transaction.From_account = fromAccount.Account_number
	transaction.To_account = toAccount.Account_number
	transaction.Amount = transactionDto.Amount
	dbc := db.Create(&transaction)
	if dbc.Error != nil {
		c.JSON(400, gin.H{"error": dbc.Error.Error()})
		return
	}
	c.JSON(201, gin.H{
		"ID":                   transaction.ID,
		"from_account":         transaction.From_account,
		"to_account":           transaction.To_account,
		"amount":               transaction.Amount,
		"to_account_balance":   toAccount.Balance,
		"from_account_balance": fromAccount.Balance,
	})
}
