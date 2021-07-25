package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"text/template"
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
	ctrl.router.HandleFunc("/api/v1/{id}", GetHello).Methods("GET")
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
func GetHello(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	buff := &bytes.Buffer{}

	tmpl,_ := template.New("helloTmpl").Parse(os.Getenv("APP_MESSAGE"))
	tmpl.Execute(buff,id)
	data := &greetingResponse{
		Msg:buff.String(),
	}
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}