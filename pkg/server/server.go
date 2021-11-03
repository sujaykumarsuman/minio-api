package server

import (
	"fmt"
	"minio-api/pkg/controller"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/labstack/gommon/log"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

func StartServer(endpoint, accessKeyID, secretAccessKey string) {
	apiServer := controller.GetMinioClient(endpoint, accessKeyID, secretAccessKey)
	HOST := "localhost"
	PORT := 8080
	mux := chi.NewMux()
	mux.Use(middleware.Logger)
	apiServer.BindRequest(mux)
	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", HOST, PORT),
		Handler: instrumentHTTPHandlerWithTracing(mux),
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server unable to start")
			log.Error(err)
		}
	}()

	fmt.Printf("Starting server at %v\n", httpServer.Addr)

	serverCloseSignal := make(chan os.Signal, 1)
	signal.Notify(serverCloseSignal, os.Interrupt)
	<-serverCloseSignal

}

func instrumentHTTPHandlerWithTracing(handler http.Handler) http.Handler {
	return nethttp.Middleware(opentracing.GlobalTracer(), handler, nethttp.OperationNameFunc(func(r *http.Request) string {
		return fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	}))
}
