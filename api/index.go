package api

import (
	"log"
	"net/http"
	"net/http/httputil"
	neturl "net/url"
)

var (
	target *neturl.URL
	proxy  *httputil.ReverseProxy
)

func init() {
	target, _ = neturl.Parse("https://www.dlsite.com")
	proxy = httputil.NewSingleHostReverseProxy(target)
}

func Proxy(w http.ResponseWriter, r *http.Request) {
	r.Host = target.Host
	proxy.ServeHTTP(w, r)
}

func Run() {
	http.HandleFunc("/", Proxy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
