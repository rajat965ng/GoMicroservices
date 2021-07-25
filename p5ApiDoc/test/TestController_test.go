package test

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"p5ApiDoc/controller"
	"strings"
	"testing"
)

func TestGetHelloHandler(t *testing.T)  {
	req := httptest.NewRequest("GET","/api/v1/happy",nil)
	w := httptest.NewRecorder()


	router := mux.NewRouter()
	router.HandleFunc("/api/v1/{id}",controller.GetHello)
	router.ServeHTTP(w,req)


	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println(w.Body.String())
	// Check the response body is what we expect.
	expected := `{"msg":"Hello, happy !!!"}`

	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}