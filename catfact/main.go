package main

import (
	"log"
)

func main() {
	svc := NewLoggingService(NewCatFactService("https://catfact.ninja/fact"))

	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":8080"))
}
