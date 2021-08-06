package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"p7ElasticSearch/domain"
	"p7ElasticSearch/repository"
	"testing"
)

func TestIndexCreation(t *testing.T) {
	conn := repository.NewConnection()
	resp, err := conn.CreateIndex("mycompany")
	if err != nil {
		fmt.Printf("Error creating index: %s\n", err.Error())
	}
	assert.NotNil(t, resp)
	fmt.Println(resp)
}

func TestDocumentCreation(t *testing.T) {
	conn := repository.NewConnection()
	emp := &domain.Employee{
		Id:          10,
		Name:        "Test User",
		Age:         25,
		Designation: "Engineer",
	}
	resp, err := conn.CreateRecord(*emp, "mycompany")
	if err != nil {
		fmt.Printf("Error creating index: %s\n", err.Error())
	}
	assert.NotNil(t, resp)
	fmt.Println(resp)
}

func TestDocumentSearch(t *testing.T) {
	conn := repository.NewConnection()
	var mapResp map[string]interface{}
	resp, err := conn.GetRecord("10", "mycompany")
	if err != nil {
		fmt.Printf("Error creating index: %s\n", err.Error())
	}
	assert.NotNil(t, resp)
	fmt.Println(resp)
	json.NewDecoder(resp.Body).Decode(&mapResp)
	fmt.Println(mapResp["hits"])
	//assert.EqualValues(t,"Test User",resp.)
}
