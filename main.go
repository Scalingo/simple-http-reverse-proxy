package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	origin, err := url.Parse(os.Getenv("ORIGIN_URL"))
	if err != nil {
		log.Println("Fail to parse URL", os.Getenv("ORIGIN_URL"))
		os.Exit(-1)
	}
	r := httputil.NewSingleHostReverseProxy(origin)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
