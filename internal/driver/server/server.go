package server

import (
	"context"
	"errors"
	"gomulticache/internal/adaptor/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const shutdownTime = 10

const v1 = "/api/v1/"

type HTTPServer struct {
	*http.Server
}

func NewHTTPServer(
	sample *controller.Sample,
) *HTTPServer {
	mux := http.NewServeMux()

	mux.Handle(strings.Join([]string{v1, controller.SamplePath, "/"}, ""), http.HandlerFunc(sample.Handle))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &HTTPServer{
		Server: s,
	}
}

func (s *HTTPServer) Run() {
	go func() {
		log.Printf("Listening on %s", s.Addr)
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}

	log.Println("HTTPServer shutdown")
}
