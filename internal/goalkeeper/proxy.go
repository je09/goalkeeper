package goalkeeper

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) (*httputil.ReverseProxy, error) {
	addr, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(addr), nil
}

func RequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := processRequest(w, r)
		if status != 200 {
			http.Error(w, "Access Forbidden by Goalkeeper", status)
			return
		}
		proxy.ServeHTTP(w, r)
	}
}
