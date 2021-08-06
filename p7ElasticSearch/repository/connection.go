package repository

import (
	"context"
	"encoding/json"
	"fmt"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"p7ElasticSearch/domain"
	"strings"
)

type connection struct {
	client *es7.Client
}

func NewConnection() *connection {
	client, err := es7.NewDefaultClient()
	if err != nil {
		fmt.Printf("Error in creating elasticsearch client: %s\n", err.Error())
	}

	fmt.Println(client.Info)

	return &connection{
		client: client,
	}
}

func (conn *connection) CreateIndex(name string) (*esapi.Response, error) {
	req := esapi.IndicesCreateRequest{
		Index: name,
	}

	return req.Do(context.Background(), conn.client)
}

func (conn *connection) CreateRecord(emp domain.Employee, index string) (*esapi.Response, error) {
	buff, _ := json.Marshal(emp)
	req := esapi.CreateRequest{
		Index:      index,
		DocumentID: string(emp.Id),
		Body:       strings.NewReader(string(buff)),
	}

	return req.Do(context.Background(), conn.client)
}

func (conn *connection) GetRecord(id string, index string) (*esapi.Response, error) {
	res, err := conn.client.Search(
		conn.client.Search.WithContext(context.Background()),
		conn.client.Search.WithIndex(index),
		conn.client.Search.WithBody(strings.NewReader(`{"query": {"match" : {"id": `+id+`}}}`)),
		conn.client.Search.WithTrackTotalHits(true),
	)
	return res, err
}
