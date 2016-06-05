package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var port = os.Getenv("PORT")
var host = os.Getenv("HOST")

func main() {
	log.SetFlags(0)
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   host,
	})
	log.Println("listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, proxy))
}
