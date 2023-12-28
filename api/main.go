package main

import (
	"github.com/ayhanks/ticketsystem/routes"
)

func main() {
	router := routes.Router()

	router.Run(":8000")
}