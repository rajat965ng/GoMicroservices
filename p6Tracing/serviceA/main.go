package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"time"
)

func initConfig() (opentracing.Tracer, io.Closer) {

	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Printf("Error in creating jeager config: %s\n", err.Error())
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Printf("Error creating tracer: %s\n", err.Error())
	}

	return tracer, closer
}

func doGET(span opentracing.Span, url string) (*http.Response, error) {

	httpClient := &http.Client{}
	httpReq, _ := http.NewRequest("GET", url, nil)
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(httpReq.Header))

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading file: %s\n", err.Error())
	}

	tracer, closer := initConfig()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	router := mux.NewRouter()

	router.HandleFunc("/svc/a/api/{id}", func(writer http.ResponseWriter, request *http.Request) {
		span := opentracing.StartSpan("Calling-Svc-B")
		span.LogFields(log.String("id", "1001"))
		defer span.Finish()

		id := mux.Vars(request)["id"]
		fmt.Printf("Id: %s\n", id)

		url := fmt.Sprintf("http://localhost:9090/svc/b/api/%s", id)

		if resp, err := doGET(span, url); err != nil { //http.Get(url)
			fmt.Printf("Error calling Get endpoint: %s\n", err.Error())
		} else {
			buff, _ := io.ReadAll(resp.Body)
			io.WriteString(writer, string(buff))
		}

	})

	svr := http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	svr.ListenAndServe()
}
