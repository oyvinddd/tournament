package main

import (
	"log"
	"tournament/app"
)

func main() {
	log.Fatal(app.New(":8080").Run())
}
