package useraccounts

import (
	"github.com/jasonnchann24/go-banking-app/helpers"
	"github.com/jasonnchann24/go-banking-app/interfaces"
)

func updateAccount(id uint, ammount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balanace", ammount)

	defer db.Close()

}
