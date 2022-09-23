package main

import (
	"flag"
	"go-proxy/internal"
	"go-proxy/internal/proxy"
	"go-proxy/testdata"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	flag.IntVar(&internal.Port, "port", internal.Port, "Proxy web server port")
	flag.StringVar(&internal.Domain, "domain", internal.Domain, "Proxy domain")
	flag.Parse()
}

func main() {
	px := proxy.New()

	for _, m := range testdata.Modules {
		_ = px.Registry.Register(m)
	}

	if err := px.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
