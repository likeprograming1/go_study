package main

import (
	"net/http"

	"main.go/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}