package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MarcosRoch4/prjdini/helpers"
	"github.com/MarcosRoch4/prjdini/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}
type Register struct {
	Username string
	Email    string
	Password string
}
type ErrResponse struct {
	Message string
}

func register(w http.ResponseWriter, r *http.Request) {
	// deixa o corpo preparado
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandlerErr(err)

	// manipula o Login
	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandlerErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	fmt.Println("Mensagem:", register["message"])

	// Prepara a resposta
	if register["message"] == "all is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		// Retorna o erro
		resp := ErrResponse{"Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	// deixa o corpo preparado
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandlerErr(err)

	// manipula o Login
	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandlerErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Prepara a resposta
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		// Retorna o erro
		resp := ErrResponse{"Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}

}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println(("App is working on port :8888"))
	log.Fatal(http.ListenAndServe(":8888", router))
}
