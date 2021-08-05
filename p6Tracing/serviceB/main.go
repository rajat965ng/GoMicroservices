package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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

func getSpanByHeaders(header http.Header) opentracing.Span {
	opsName := "Called-Svc-B"

	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	if err != nil {
		fmt.Printf("Error creating span contxt: %s\n", err.Error())
	}

	return opentracing.StartSpan(opsName, ext.RPCServerOption(spanCtx))
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading file: %s\n", err.Error())
	}

	tracer, closer := initConfig()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	router := mux.NewRouter()

	router.HandleFunc("/svc/b/api/{id}", func(writer http.ResponseWriter, request *http.Request) {
		span := getSpanByHeaders(request.Header)
		defer span.Finish()

		id := mux.Vars(request)["id"]
		fmt.Printf("Id: %s", id)

		io.WriteString(writer, fmt.Sprintf("Hello, %s !!!", id))
	})

	svr := http.Server{
		Handler:      router,
		Addr:         ":9090",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	svr.ListenAndServe()
}
