package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func StartServer(port string, done chan os.Signal) {
	router := mux.NewRouter()

	// Define the REST endpoints
	router.HandleFunc("/hello-world", HelloWorld).Methods("GET")

	fmt.Printf("\n🚀 REST API Server ready at http://localhost:%s/ \n\n", port)

	// Start the REST API server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("rest-api -> Server stopped: ", err)
		}
	}()

	<-done // wait for an interrupt signal

	fmt.Println("rest-api -> Received an interrupt signal, shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("rest-api -> Server shutdown error:", err)
	}

	fmt.Println("REST API Server shut down gracefully!")
}
