package forkort

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func ShortenLink(url string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	resp, err := client.Post("https://forkort.dk/api/shorten", "application/json", strings.NewReader(url))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", err
	}
	return string(body), nil
}

func UnshortenLink(token string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	resp, err := client.Get("https://forkort.dk/api/unshorten/" + token)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", err
	}
	return string(body), nil
}
