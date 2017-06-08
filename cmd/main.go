package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/nicklanng/gecho/pkg"
)

func main() {
	service := getConfig()

	http.HandleFunc("/", pkg.Handler(service))
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getConfig() (service string) {
	servicePtr := flag.String("service", "gecho", "A service name to be returned in the responses")
	flag.Parse()

	service = *servicePtr

	return
}
