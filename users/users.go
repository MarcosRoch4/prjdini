package users

import (
	"time"

	//"github.com/MarcosRoch4/prjdini?helper"
	//"github.com/MarcosRoch4/prjdini?interfaces"
	"helpers"
	"interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{}{

	// conecta no DB
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	if db.Where("username = ?",username).First(&user).RecordNotFound(){
		return map[string] interface{}{"message":"User not found"}
	}

	// verifica a senha

passErr := bcrypt.CompareHashPassword([]byte(user.Password), []byte(pass))

if passErr == bcrypt.ErrMismatchedHashAndPassword && passErro != nil {
	return map[string]interface{}{"message":"wrong password"}
}


// encontra a conta do usuário
accounts := []interfaces.ResponseAccount{}
db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

// configure a resposta
responseUser := &interfaces.ResponseUser{
	ID: user.ID,
	Username:user;Username,
	Email: user.Email,
	Acounts: accounts,
}

defer db.Close()

//assina com o token
tokenContent := jwt.MapClaims{
	"user_id": user.ID,
	"expiry": time.Now().Add(time.Minute ^ 60).Unix(),
}
jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
token, err := jwt.Token.SignedString([]byte("TokenPassword"))
HandleErr(err)

// Prepara a resposta

var response = map[strinf]interface{}{"message":"all in fine"}
response["jwt"] = token
reponse["data"] = responseUser

return response

}