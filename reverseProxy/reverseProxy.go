package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
	"net/url"
)


func main() {
	url, err := url.Parse("http://localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	http.Handle("/proxy", &ProxyHandler{proxy})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ph.p.ServeHTTP(w, r)
}
