package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MarcosRoch4/prjdini/helpers"
	"github.com/MarcosRoch4/prjdini/interfaces"
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

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandlerErr(err)
	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "all in fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		// Retorna o erro
		resp := interfaces.ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	// deixa o corpo preparado
	body := readBody(r)

	// manipula o Login
	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandlerErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Prepara a resposta
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	// deixa o corpo preparado
	body := readBody(r)

	// manipula o Login
	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandlerErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	fmt.Println("Mensagem:", register["message"])

	// Prepara a resposta
	apiResponse(register, w)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)

	apiResponse(user, w)

}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	fmt.Println(("App is working on port :8888"))
	log.Fatal(http.ListenAndServe(":8888", router))
}
