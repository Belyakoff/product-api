package main

import (
	"github.com/Belyakoff/product-api/handlers"
	"net/http"
	"os" 
	"os/signal"
 	"log"
 	"time"
 	"context"	
) 


func main(){

	l  := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
		 		l.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Received terminate, graceful shutdown:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}