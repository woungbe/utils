package utils

import (
	"fmt"
	"testing"
)

func TestInsertQuery(t *testing.T) {

	type User struct {
		ID       int
		Username string
		Email    string
	}

	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	insertQuery := GenerateInsertQuery("users", user)
	fmt.Println(insertQuery)
}

func TestInsertOrUpdate(t *testing.T) {

	type DepositMaster struct {
		Txid                string `json:"txid" db:"txid"`
		CoboID              string `json:"coboID" db:"coboID"`
		Status              string `json:"status" db:"status"`
		Coin                string `json:"coin" db:"coin"`
		Network             string `json:"network" db:"network"`
		FromAddress         string `json:"fromAddress" db:"fromAddress"`
		ToAddress           string `json:"toAddress" db:"toAddress"`
		ConfirmedNumber     int    `json:"confirmedNumber" db:"confirmedNumber"`
		ConfirmingThreshold int    `json:"confirmingThreshold" db:"confirmingThreshold"`
		TransactionType     int    `json:"transactionType" db:"transactionType"`
		CreatedTime         int    `json:"createdTime" db:"createdTime"`
		UpdatedTime         int    `json:"updatedTime" db:"updatedTime"`
		GasPrice            int    `json:"gasPrice" db:"gasPrice"`
		GasLimit            int    `json:"gasLimit" db:"gasLimit"`
		FeeUsed             int    `json:"feeUsed" db:"feeUsed"`
	}

	user := DepositMaster{
		Txid:                "Txid",
		CoboID:              "CoboID",
		Status:              "500",
		Coin:                "Coin",
		Network:             "Network",
		FromAddress:         "0x123123",
		ToAddress:           "0x123124",
		ConfirmedNumber:     7,
		ConfirmingThreshold: 64,
		TransactionType:     900,
		CreatedTime:         1613912309123,
		UpdatedTime:         16192381923891,
		GasPrice:            1000000,
		GasLimit:            1000000,
		FeeUsed:             200000000,
	}

	// query := GenerateInsertOrUpdateQuery("users", user)
	// /// query := GenerateInsertOrUpdateQuery("users", user, "Username = VALUES(Username)")
	query := GenerateInsertOrUpdateQuery("users", user, []string{"ConfirmedNumber"})
	fmt.Println("query :", query)
}
