package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gregjones/httpcache/memcache"
	"github.com/jackc/pgx/v4"
	"net/http"
	"time"
)

type Abc struct {
	Nconst              string   `json:"nconst"`
	Primary_name        string   `json:"primary_name"`
	Birth_year          string   `json:"birth_year"`
	Death_year          string   `json:"death_year"`
	Primary_professions []string `json:"primary_professions"`
	Known_for_titles    []string `json:"known_for_titles"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/name/{id}", getNameByIdHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("error while starting server: %s\n", err.Error())
	}
}

func getNameByIdHandler(resp http.ResponseWriter, req *http.Request) {
	identity := mux.Vars(req)["id"]

	cache := memcache.New("localhost:11211")
	cache.Timeout = 100 * time.Second
	cache.MaxIdleConns = 50

	fmt.Printf("cache conn is %s\n", cache)

	res := Abc{}

	if data, ok := cache.Get(identity); ok == true {
		fmt.Println("MemCache start !!")
		fmt.Printf("data is %s\n", data)
		if err := json.Unmarshal(data,&res);err!=nil {
			fmt.Printf("Error in unmarshalling data to payload: %s\n",err.Error())
		}
		fmt.Printf("ok is %s\n", ok)
		fmt.Printf("res is %s\n", res)
	} else {
		fmt.Println("PgSql start !!")
		pgConn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/dbname?sslmode=disable")
		if err != nil {
			fmt.Printf("Error creating  postgres connnection: %s\n", err.Error())
		}
		defer pgConn.Close(context.Background())
		query := `SELECT nconst, primary_name, birth_year, death_year,primary_professions FROM "names" WHERE nconst = $1`
		if err := pgConn.QueryRow(context.Background(), query, identity).Scan(&res.Nconst, &res.Primary_name, &res.Birth_year, &res.Death_year,&res.Primary_professions); err != nil {
			fmt.Printf("Error querying: %s\n", err.Error())
		} else {
			if data,err := json.Marshal(res);err!=nil {
				fmt.Printf("Error in marshalling  data: %s\n",err.Error())
			}else {
				cache.Set(res.Nconst, data)
			}
		}

	}

	resp.WriteHeader(http.StatusOK)

	fmt.Printf("response payload is: %s\n", res)
	json.NewEncoder(resp).Encode(res)

}
