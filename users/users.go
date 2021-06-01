package users

import (
	"fmt"
	"time"

	"github.com/MarcosRoch4/prjdini/helpers"
	"github.com/MarcosRoch4/prjdini/interfaces"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func prepareToken(user *interfaces.User) string {
	//conecta com o token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandlerErr(err)

	return token

}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount, withToken bool) map[string]interface{} {
	// configure a resposta
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	// Prepara a resposta

	var response = map[string]interface{}{"message": "all is fine"}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}

	response["data"] = responseUser

	fmt.Println("message:", responseUser)

	return response

}

func Login(username string, pass string) map[string]interface{} {

	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if valid {
		// conecta no DB
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		// verifica a senha

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "wrong password"}
		}

		// encontra a conta do usuário
		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		defer db.Close()

		var response = prepareResponse(user, accounts, true)

		return response
	} else {
		return map[string]interface{}{"message": "not Valid values"}
	}
}

func Register(username string, email string, pass string) map[string]interface{} {
	fmt.Println("Chegou no Register do pacote users")
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	fmt.Println(valid)
	if valid {
		fmt.Println("Os dados estão válidos")
		db := helpers.ConnectDB()
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(username + "'s" + " account"),
			Balance: 0, UserId: user.ID}

		db.Create(&account)

		defer db.Close()

		accounts := []interfaces.ResponseAccount{}
		respAccount := interfaces.ResponseAccount{ID: account.ID, Name: account.Name, Balance: int(account.Balance)}
		accounts = append(accounts, respAccount) // res []interfaces.ResponseAccount) map[string]interfaces{}
		var response = prepareResponse(user, accounts, true)

		return response

	} else {
		return map[string]interface{}{"message": "not Valid values"}
	}
}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	fmt.Println("é válido?", isValid)
	if isValid {
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		defer db.Close()

		var response = prepareResponse(user, accounts, false)

		return response

	} else {
		return map[string]interface{}{"message": "not Valid token"}
	}
}
