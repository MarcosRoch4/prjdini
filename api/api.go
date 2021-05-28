package api


type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request){
	// deixa o corpo preparado
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	// manipula o Login
	var formattedBody Login
	err = json.Unmarshal(body,&formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Prepara a resposta
	if login["message"] == "all in fine"{
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else{
		// Retorna o erro
		resp := ErrResponse{"Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}

}

func StartApi(){
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println(("App is working on port :8888"))
	log.Fatal(http.ListenAndServe(":8888",router))
}