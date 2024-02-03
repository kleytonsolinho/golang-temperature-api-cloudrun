package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kleytonsolinho/golang-temperature-api-cloudrun/configs"
	"github.com/kleytonsolinho/golang-temperature-api-cloudrun/internal/infra/webserver/handlers"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("WEATHER_API_KEY", configs.WEATHER_API_KEY))

	r.Get("/{cep}", handlers.GetCepHandler)

	http.ListenAndServe(":8080", r)
}
