package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
)

type Cep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type Temperature struct {
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

type ResponseBody struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetCepHandler(w http.ResponseWriter, r *http.Request) {
	cepParams := chi.URLParam(r, "cep")
	if len(cepParams) != 8 || cepParams == "00000000" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid zipcode")
		return
	}

	apiKey, ok := r.Context().Value("WEATHER_API_KEY").(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Erro ao obter a chave da API de tempo.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	chLocale := make(chan string)
	chTempurature := make(chan float64)

	go getCepViaCEP(cepParams, chLocale)
	go getTemperature(chTempurature, chLocale, apiKey)

	for {
		select {
		case temperature := <-chTempurature:
			if temperature == 0 {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode("can not find temperature")
				return
			}

			response := getTemperatureFandK(temperature)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			return
		case <-ctx.Done():
			w.WriteHeader(http.StatusRequestTimeout)
			json.NewEncoder(w).Encode("timeout exceeded")
			return
		}
	}
}

func getCepViaCEP(cepParams string, chLocale chan string) {
	req, err := http.NewRequest("GET", "http://viacep.com.br/ws/"+cepParams+"/json/", nil)
	if err != nil {
		chLocale <- ""
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		chLocale <- ""
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		chLocale <- ""
		panic(err)
	}

	var result Cep
	err = json.Unmarshal(body, &result)
	if err != nil {
		chLocale <- ""
		panic(err)
	}

	log.Printf("Response ViaCEP: %v", result)
	chLocale <- result.Localidade
}

func getTemperature(chTemperature chan float64, chLocale chan string, apiKey string) {
	locale := <-chLocale
	if locale == "" {
		return
	}

	escapedLocale := url.QueryEscape(locale)

	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json?q="+escapedLocale+"&key="+apiKey, nil)
	if err != nil {
		chTemperature <- 0
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		chTemperature <- 0
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		chTemperature <- 0
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Erro na resposta HTTP. Código: %d\n", res.StatusCode)
		fmt.Println("Corpo da resposta:", string(body))
		chTemperature <- 0
		return
	}

	var resultTemperature Temperature
	err = json.Unmarshal(body, &resultTemperature)
	if err != nil {
		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
		chTemperature <- 0
		return
	}

	log.Printf("Response Temperature C: %v", resultTemperature.Current.TempC)
	log.Printf("Response Temperature F: %v", resultTemperature.Current.TempF)
	chTemperature <- resultTemperature.Current.TempC
}

func getTemperatureFandK(tempC float64) ResponseBody {
	tempF := (tempC * 1.8) + 32
	tempK := tempC + 273

	return ResponseBody{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}
}
