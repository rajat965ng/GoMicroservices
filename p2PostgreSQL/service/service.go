package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rajat965ng/p2PostgreSQL/repository"
	"net/http"
)

type Service struct {
	psql *repository.PostgreSQL
}

func PsqlService(psql *repository.PostgreSQL) *Service  {
	return &Service{
		psql:psql,
	}
}

func (svc *Service) GetNameById(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	fmt.Printf("Query name for id: %s\n",id)
	name := svc.psql.FindUserById(id)
	json.NewEncoder(writer).Encode(name)
}