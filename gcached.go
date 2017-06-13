package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// the global buckets list
// var buckets map[string]Bucket
var buckets = map[string]Bucket{}

func main() {

	// the global buckets list
	buckets = map[string]Bucket{}

	fmt.Println("gcached starting up...")

	stop := make(chan os.Signal)
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":3000"
	}

	signal.Notify(stop, os.Interrupt)

	router := NewRouter()
	h := &http.Server{Addr: addr, Handler: router}

	go func() {
		fmt.Println(fmt.Sprintf("Listening on %s. CTRL-C to shutdown.\n", addr))
		if err := h.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop

	fmt.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	h.Shutdown(ctx)
}
