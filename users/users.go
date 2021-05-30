package users

import (
	"time"

	"github.com/MarcosRoch4/prjdini/helpers"
	"github.com/MarcosRoch4/prjdini/interfaces"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {

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

	// configure a resposta
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	defer db.Close()

	//assina com o token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandlerErr(err)

	// Prepara a resposta

	var response = map[string]interface{}{"message": "all in fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response

}
