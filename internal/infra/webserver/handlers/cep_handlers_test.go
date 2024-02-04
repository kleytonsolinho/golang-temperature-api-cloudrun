package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetCepHandlerInvalid1(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/cep/{cep}", GetCepHandler)

	req, err := http.NewRequest("GET", "/cep/00000000", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Body.String() != "\"invalid zipcode\"\n" {
		t.Errorf("Esperado 'invalid zipcode', mas obteve '%s'", w.Body.String())
	}
}

func TestGetCepHandlerInvalid2(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/cep/{cep}", GetCepHandler)

	req, err := http.NewRequest("GET", "/cep/2525526", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Body.String() != "\"invalid zipcode\"\n" {
		t.Errorf("Esperado 'invalid zipcode', mas obteve '%s'", w.Body.String())
	}
}

func TestGetCepHandlerSucess(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/cep/{cep}", GetCepHandler)

	req, err := http.NewRequest("GET", "/cep/01001000", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado status HTTP 200, mas obteve %d", w.Code)
	}
}

func TestGetCepViaCEP(t *testing.T) {
	result, err := getCepViaCEP("25255260")
	if err != nil {
		t.Fatal(err)
	}

	if result.Cep != "25255-260" {
		t.Errorf("Esperado CEP '25255-260', mas obteve '%s'", result.Cep)
	}
}

func TestGetTemperature(t *testing.T) {
	result, err := getTemperature("Duque de Caxias")
	if err != nil {
		t.Fatal(err)
	}

	if result.Current.TempC == 0 {
		t.Error("Esperado um valor diferente de 0 para a temperatura, mas obteve 0")
	}
}

func TestGetTemperatureFandK(t *testing.T) {
	result := getTemperatureFandK(30)
	if result.TempC != 30 {
		t.Errorf("Esperado 30, mas obteve %f", result.TempC)
	}
	if result.TempF != 86 {
		t.Errorf("Esperado 86, mas obteve %f", result.TempF)
	}
	if result.TempK != 303 {
		t.Errorf("Esperado 303, mas obteve %f", result.TempK)
	}
}
