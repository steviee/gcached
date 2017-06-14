package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/steviee/gcached/service"
	"github.com/steviee/gcached/store"
)

func main() {

	fmt.Println("gcached starting up...")

	back := store.StartBackgroundDump()

	stop := make(chan os.Signal)
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":3000"
	}

	signal.Notify(stop, os.Interrupt)

	router := service.NewRouter()
	h := &http.Server{Addr: addr, Handler: router}

	go func() {
		fmt.Println(fmt.Sprintf("Listening on %s. CTRL-C to shutdown.\n", addr))
		if err := h.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop // stop the server

	fmt.Println("\nShutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	back <- true

	// dump the data (finally)
	store.DumpToDisk()

	h.Shutdown(ctx)
	fmt.Println("\nShutdown complete...")

}
