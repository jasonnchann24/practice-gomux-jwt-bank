package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jasonnchann24/go-banking-app/helpers"
	"github.com/jasonnchann24/go-banking-app/interfaces"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {

	db := helpers.ConnectDB()
	user := &interfaces.User{}

	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong Password"}
	}

	// find account for the user
	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

	// response
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	defer db.Close()

	// sign token

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	// Prepare response
	var response = map[string]interface{}{"message": "All is good"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
