package main

import "template-api-golang/config"

func main() {
	// Carregar container
	context, container, err := config.LoadContainer()

	if err != nil {
		panic(err)
	}

	// Adaptador HTTP
	// Carrega adpatador externo do HTTP
	router := config.LoadRouter(context)

	// Carrega servidor HTTP
	config.LoadHTTP(context, container, router)
}
