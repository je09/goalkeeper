package goalkeeper

import (
	"Goalkeeper/pkg/config"
	"fmt"
	"log"
	"net/http"
)

func Start() {
	log.Printf("Starting Goalkeeper")
	cnf := config.Read()
	err := InitWaf(cnf)
	if err != nil {
		panic(fmt.Errorf("Fatal error coroza setup: %w \n", err))
	}
	handle(cnf)
}

func handle(cnf config.Config) {
	proxy, err := NewReverseProxy(cnf.Proxy.AppAddress)
	if err != nil {
		panic(fmt.Errorf("Fatal error proxy setting: %w \n", err))
	}
	http.HandleFunc("/", RequestHandler(proxy))
	log.Printf("Listening on port: %d", cnf.Proxy.WafPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cnf.Proxy.WafPort), nil))
}
