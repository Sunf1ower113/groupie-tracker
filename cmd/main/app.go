package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	artist "groupie-tracker/internal/artist"
	config "groupie-tracker/internal/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	handler := artist.NewHandler()
	log.Println("register routes")

	start(handler.Mux)
}

func start(router *http.ServeMux) {
	log.Println("Start the application...")
	cfg, err := config.LoadConfiguration("config.json")
	if err != nil {
		log.Fatal(err)
	}
	listner, err := net.Listen(cfg.Listen.Protocol, fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is listening port %s:%s\n", cfg.Listen.BindIp, cfg.Listen.Port)
	log.Fatal(server.Serve(listner))
}
