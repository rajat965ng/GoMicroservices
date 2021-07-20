package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rajat965ng/p2PostgreSQL/repository"
	"github.com/rajat965ng/p2PostgreSQL/service"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Starting application !!!\n")
	conn, err := repository.NewConection()
	if err != nil {
		fmt.Printf("Error in creating new connection: %s\n", err.Error())
	} else {
		defer conn.Close()
	}

	svc := service.PsqlService(conn)
	router := mux.NewRouter()

	router.HandleFunc("/names/psql/{id}", svc.GetNameById)

	server := http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error in starting server: %s\n", err)
	}
}
