# gecho
A simple echo server, a la httpbin

[![Build Status](https://travis-ci.org/nicklanng/gecho.svg?branch=master)](https://travis-ci.org/nicklanng/gecho)

## Usage

Command-line arguments
```sh
  -cert string
        The certificate for server SSL (default "server.crt")
  -key string
        The key for server SSL (default "server.key")
  -port int
        Which port to listen on (default 8080)
  -service string
        A service name to be returned in the responses (default "gecho")
  -ssl
        Use SSL. Recommend also set the cert and key args.
```
