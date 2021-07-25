package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
	"net/http"
	"p5ApiDoc/controller"
	"time"
	_ "p5ApiDoc/docs"  // docs is generated by Swag CLI, you have to import it.
)

// @title Test API
// @version 1.0
// @description This is a sample service for managing Ids
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email rajatnigam89@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main()  {
	if err := godotenv.Load();  err!=nil {
		fmt.Printf("Error is loading configuration: %s\n",err.Error())
	}
	router := mux.NewRouter()
	ctrl,_ := controller.NewController(router)
	ctrl.Register()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server := http.Server{
		Handler:router,
		Addr:":8080",
		ReadTimeout:15*time.Second,
		WriteTimeout:15*time.Second,
	}
	server.ListenAndServe()
}