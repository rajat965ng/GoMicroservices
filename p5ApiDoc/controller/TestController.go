package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type controller struct {
	router *mux.Router
}


type greetingResponse struct {
	Msg string `json:"msg"`
}

func NewController(router *mux.Router) (*controller,error)  {
	return &controller{
		router:router,
	}, nil
}

func (ctrl *controller) Register()  {
	ctrl.router.HandleFunc("/api/v1/{id}", getHello).Methods("GET")
}

// GetHello godoc
// @Summary Hello
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "Api Id"
// @Success 200 {object} greetingResponse
// @Router /api/v1/{id} [get]
func getHello(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	data := &greetingResponse{
		Msg:fmt.Sprintf(os.Getenv("APP_MESSAGE"),id),
	}
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}