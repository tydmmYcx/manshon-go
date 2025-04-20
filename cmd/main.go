package main

import (
	"fmt"
	"net/http"
	"time"

	"manshon-go/internal/initialize"
)

func init() {

}

func main() {

	fmt.Println("Start Manshon-go!")

	Router := initialize.Router()
	s := &http.Server{
		Addr:           ":8090",
		Handler:        Router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
