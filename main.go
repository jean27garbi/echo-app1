package main

import (
	"github.com/jean27garbi/echo-app1/routes"
)

func main() {

	router := routes.UserRouter()
	router.Logger.Fatal(router.Start(":3000"))
}
