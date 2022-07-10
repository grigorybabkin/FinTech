package main

import (
	"FinTech/storage"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type URL struct {
	FullURL  string `json:"full_url"`
	ShortURL string `json:"short_url"`
}

const (
	DB     = 1
	Memory = 0
)

var (
	urlMem  = storage.NewURLMemory()
	Storage int
)

func main() {

	cmdCorrect := false
	var cmd string

	http.HandleFunc("/", handler)
	//fmt.Fprintln(os.Stdout,"Choose way to store urls")

	for {
		fmt.Fprintln(os.Stdout, "enter one of these variants:")
		fmt.Fprintln(os.Stdout, "->   DB or 1")
		fmt.Fprintln(os.Stdout, "->   M  or 0")

		fmt.Fscan(os.Stdin, &cmd)

		switch cmd {
		case "DB", "1":
			Storage = 1
			cmdCorrect = true
			fmt.Fprintln(os.Stdout, "DB mode chosen")
		case "M", "0":
			Storage = 0
			cmdCorrect = true
			fmt.Fprintln(os.Stdout, "Memory mode chosen")
		}
		if cmdCorrect {
			break
		}
	}

	/*	storage := flag.Int("storage", 1, "Type of url storage")
		Storage = *storage
		switch Storage {
		case 0:
			fmt.Fprintln(os.Stdout, "URL storage is Memory")
		default:
			fmt.Fprintln(os.Stdout, "URL storage is DB")
		}
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}
