package main

import (
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:9001/test", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-istio", "a\xc5z")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("resp: %v\n", resp)
}
