package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<div><h1>Welcome to the temperature API</h1><a href='/cep/01001000'>Click Here <p style='color:red;'>/cep/01001000</p></a></div>"))
}
