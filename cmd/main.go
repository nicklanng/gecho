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
	service, certPath, keyPath, port, ssl := getConfig()

	listenAddress := ":" + strconv.Itoa(port)

	http.HandleFunc("/", pkg.Handler(service))
	fmt.Println("listening on", listenAddress)

	if ssl {
		log.Fatal(http.ListenAndServeTLS(listenAddress, certPath, keyPath, nil))
	} else {
		log.Fatal(http.ListenAndServe(listenAddress, nil))
	}
}

func getConfig() (service, certPath, keyPath string, port int, ssl bool) {
	servicePtr := flag.String("service", "gecho", "A service name to be returned in the responses")
	certPathPtr := flag.String("cert", "server.crt", "The certificate for server SSL")
	keyPathPtr := flag.String("key", "server.key", "The key for server SSL")
	portPtr := flag.Int("port", 8080, "Which port to listen on")
	sslPtr := flag.Bool("ssl", false, "Use SSL. Recommend also set the cert and key args.")

	flag.Parse()

	service = *servicePtr
	certPath = *certPathPtr
	keyPath = *keyPathPtr
	port = *portPtr
	ssl = *sslPtr

	return
}
