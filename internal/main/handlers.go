package main

import (
	"FinTech/pkg"
	"FinTech/storage"
	"net/http"
	"net/url"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GETHandler(w, r)
	case http.MethodPost:
		POSTHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GETHandler(w http.ResponseWriter, r *http.Request) {

	shortURL := r.URL.Path[1:]
	//:= r.URL.Query().Get("short_url")
	var fullURL string
	switch Storage {
	case DB:
		if storage.RowExistsDB("SELECT * FROM urls WHERE short_url=$1", shortURL) {
			fullURL = storage.GetURLDB("SELECT full_url FROM urls WHERE short_url = $1", shortURL)
		}
		//else ТАКОЙ ССЫЛКИ НЕТ
	case Memory:
		fullURL, _ = urlMem.GetURLMem(shortURL)
	}

	w.Write([]byte(fullURL))
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {

	fullURL := r.URL.Query().Get("full_url")

	RawFullURL, _ := url.Parse(fullURL)

	RawQuery := RawFullURL.RawQuery

	if len(RawQuery) != 0 {

		values := url.Values{}
		params := regexp.MustCompile("[=&]{1}").Split(RawQuery, -1)

		for idx := 0; idx <= len(params)/2; idx += 2 {
			values.Add(params[idx], params[idx+1])
		}

		RawFullURL.RawQuery = values.Encode()
	}

	fullURL = RawFullURL.String()
	shortURL := pkg.RandShortURL(10)

	switch Storage {
	case DB:
		if !storage.RowExistsDB("SELECT * FROM urls WHERE full_url=$1", fullURL) {
			storage.AddURLsDB(fullURL, shortURL)
		} else {
			shortURL = storage.GetURLDB("SELECT short_url FROM urls WHERE full_url = $1", fullURL)
		}
	case Memory:
		urlMem.AddURLsMem(shortURL, fullURL, shortURL)
	}

	w.WriteHeader(http.StatusOK)
	//fullURLByte, _ := json.Marshal(fullURL)
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(fullURL))
	w.Write([]byte(shortURL))
}
