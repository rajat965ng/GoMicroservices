package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rajat965ng/p3Gorm/service"
	"net/http"
)

type nameController struct {
	router *mux.Router
	nameSvc *service.NameService
}

func NewNameController(router *mux.Router) (*nameController)  {
	nameSvc,err := service.NewNameService()
	if err!=nil {
		fmt.Printf("Error in creating name service: %s\n",err.Error())
	}
	return &nameController{
		router:router,
		nameSvc:nameSvc,
	}
}

func (nc *nameController) Register()  {
	nc.router.HandleFunc("/name/gorm/{id}", nc.GetNamesById)
}

func (nc *nameController) GetNamesById(writer http.ResponseWriter, request *http.Request)  {
	id := mux.Vars(request)["id"]
	name,_ := nc.nameSvc.FindNameService(id)
	json.NewEncoder(writer).Encode(name)
}