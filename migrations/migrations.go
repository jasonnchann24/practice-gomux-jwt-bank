package migrations

import (
	"github.com/jasonnchann24/go-banking-app/helpers"
	"github.com/jasonnchann24/go-banking-app/interfaces"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Username: "John", Email: "john@email.com"},
		{Username: "Doe", Email: "doe@email.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}

		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}

	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)

	defer db.Close()

	createAccounts()
}
