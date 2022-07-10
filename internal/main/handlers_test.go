package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/KYH_2geLYN", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GETHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPOSTHandler(t *testing.T) {

	req, err := http.NewRequest("POST", "?full_url=https://www.ozon.ru/category/elektronika-15500/",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(POSTHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
