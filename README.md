# gecho
A simple echo server, a la httpbin

[![Build Status](https://travis-ci.org/nicklanng/gecho.svg?branch=master)](https://travis-ci.org/nicklanng/gecho)

## Installation
If you have go installed, simply run: `go get github.com/nicklanng/gecho`.

Otherwise, click on the `releases` tab to find executables for Windows, Linux and OS X.

## Usage

Launch the application, configuring options can be seen below.
```sh
$ ./gecho-linux-amd64 -service=test
```

Send a request to the endpoint. You get a response showing the details of the request, including headers, query parameters and url.
```sh
$ curl -X POST -i -H x-some-header:4535 "localhost:8080/418/other/things?cool=yes"
{"service":"test","host":"localhost:8080","url":"/418/other/things?cool=yes","args":{"cool":["yes"]},"method":"POST","origin":"[::1]:63708","headers":{"Accept":["*/*"],"User-Agent":["Mozilla/5.0 (Windows; U; MSIE 9.0; WIndows NT 9.0; en-US))"],"X-Some-Header":["4535"]},"body":""}
```

You can specify the response status code by including as the first path segment. Any unknown status code will be returned as 200 OK.
```sh
$ curl localhost:8080/201 # 201 Created
$ curl localhost:8080/404 # 404 Not Found
$ curl localhost:8080/fdgdg  # 200 OK
```

### Command-line arguments
You can see this from the command by launching it with the `-h` argument.
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
