package main

import (
	"github.com/gorilla/mux"
	"github.com/rajat965ng/p3Gorm/controller"
	"net/http"
	"time"
)

func main() {
	router := &mux.Router{}
	nc := controller.NewNameController(router)
	nc.Register()
	
	server := http.Server{
		Handler:router,
		ReadTimeout: 15*time.Second,
		WriteTimeout: 15*time.Second,
		Addr: ":8080",
	}
	
	server.ListenAndServe()
}
