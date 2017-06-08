package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nicklanng/gecho/pkg"
)

func main() {
	service, port := getConfig()

	portString := strconv.Itoa(port)

	http.HandleFunc("/", pkg.Handler(service))
	fmt.Println("listening on :" + portString)
	log.Fatal(http.ListenAndServe(":"+portString, nil))
}

func getConfig() (service string, port int) {
	servicePtr := flag.String("service", "gecho", "A service name to be returned in the responses")
	portPtr := flag.Int("port", 8443, "Which port to listen on")
	flag.Parse()

	service = *servicePtr
	port = *portPtr

	return
}
