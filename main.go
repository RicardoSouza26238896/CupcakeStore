package main

import (
	"fmt"
	"log"

	"ricardosouza26238896.github.io/cupcakestore/bootstrap"
	"ricardosouza26238896.github.io/cupcakestore/config"
)

func main() {
	app := bootstrap.NewApplication()
	cfg := config.Instance()

	host := cfg.GetEnvVar("APP_HOST", "localhost")
	port := cfg.GetEnvVar("APP_PORT", "4000")
	addr := fmt.Sprintf("%s:%s", host, port)

	if cfg.GetEnvVar("DEV_MODE", "true") == "true" {
		log.Fatal(app.Listen(addr))
		return
	}

	certFile := "/etc/letsencrypt/live/ricardosouza26238896.github.io/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/ricardosouza26238896.github.io/privkey.pem"
	log.Fatal(app.ListenTLS(addr, certFile, keyFile))
}
