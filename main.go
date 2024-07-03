package main

import (
	"fmt"
	"github.com/KapilSareen/go-project/pkg/api"
)

 func main(){
    fmt.Println("Loading...")
	fmt.Println("Listening on http://localhost:8001/")
	api.Start()
}