package main

import (
	"net/http"
	"os"
 	"log"
	"github.com/Belyakoff/product-api/handlers"
) 


func main(){

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
			
	})


	http.ListenAndServe(":9090", nil)

}