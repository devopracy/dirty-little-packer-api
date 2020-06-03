package main

import (
	"log"
	"net/http"

	"github.com/munjeli/dirty-little-packer-api/api/dlpa"
)

func main() {
	http.HandleFunc("/status", dlpa.StatusHandler)
	http.HandleFunc("/build", dlpa.BuildHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
