package pkg

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type response struct {
	Service string      `json:"service"`
	Host    string      `json:"host"`
	URL     string      `json:"url"`
	Args    url.Values  `json:"args"`
	Method  string      `json:"method"`
	Origin  string      `json:"origin"`
	Header  http.Header `json:"headers"`
	Body    string      `json:"body"`
}

func (r *response) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return b, err
}
