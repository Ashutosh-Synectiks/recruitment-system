package main

import (
	"fmt"
	"log"
	"net/http"
	"recruitment_system/routes"
)

func main() {
	r := routes.SetupRouter()
	fmt.Println("start")
	log.Println(http.ListenAndServe(":8000", r))

}
