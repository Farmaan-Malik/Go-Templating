package main

import (
	"fmt"
	"net/http"
	"github.com/Farmaan-Malik/Go-Templating/pkg/handlers"
)

const portNumber = ":8080"

func main (){
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server started on port %s", portNumber)
	_ = http.ListenAndServe(portNumber,nil)
}