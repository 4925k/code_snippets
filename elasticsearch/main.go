package main

import (
	"fmt"
	"log"

	"github.com/badkaktus/gorocket"
	"github.com/elastic/go-elasticsearch/v7"
)

var (
	user     = "elastic"
	password = "password"
	es       = "http://127.0.0.1:9200"
	rcURL    = "https://chat.vairav.net/rBDgHCdKQydtxc2Si/7MuRv42dw7icn5JZaw4g5mzi6SpGTQRjDHhHbZ6QKv9BmoMa"
)

func main() {
	fmt.Println(esConnection(es, user, password))
	fmt.Println(rcConnection(rcURL))
}

func esConnection(url, user, password string) bool {
	cfg := elasticsearch.Config{
		Username:  user,
		Password:  password,
		Addresses: []string{url},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("[FATAL] create connection: %v", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Printf("[ERROR] get info: %v", err)
	}
	defer res.Body.Close()

	return !res.IsError()
}

func rcConnection(url string) bool {
	rc := gorocket.NewClient(url)

	msg := gorocket.Message{
		Text: "connection test",
	}

	rc.PostMessage(&msg)

	return false
}
