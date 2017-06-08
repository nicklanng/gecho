package pkg

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Handler(service string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := getStatusCode(r.URL)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		res := new(response)
		res.Service = service
		res.Header = r.Header
		res.URL = r.URL.String()
		res.Method = r.Method
		res.Origin = r.RemoteAddr
		res.Host = r.Host
		res.Body = string(body)
		res.Args = r.URL.Query()

		bytes, err := res.JSON()
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(statusCode)
		w.Write(bytes)
	}
}

func getStatusCode(url *url.URL) int {
	status := strings.Split(url.Path[1:], "/")[0]
	code, err := strconv.Atoi(status)
	if err != nil {
		return 200
	}
	return code
}
