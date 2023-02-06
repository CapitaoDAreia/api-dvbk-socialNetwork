package main

import (
	config "api-dvbk-socialNetwork/configuration"
	"api-dvbk-socialNetwork/internal/infraestructure/http/auth"
	"api-dvbk-socialNetwork/internal/infraestructure/http/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	auth.GenerateSecretKey()
}

func main() {
	config.LoadAmbientConfig()
	fmt.Printf("PORT=%v\n", config.PORT)

	//Generate routes to feed Server
	r := router.Generate()

	//Generate PORT valur to feed Server
	var PORT = fmt.Sprintf(":%v", config.PORT)

	fmt.Printf("Listening on PORT %v...\n", config.PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
