package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
	"time"
)

type (
	Client struct {
		name   string
		events chan *DashBoard
	}
	DashBoard struct {
		User uint
	}
)

func main() {
	app := fiber.New()
  app.Static("/", "./static")
	app.Get("/sse/validation", adaptor.HTTPHandler(handler(dashbaordHandler)))
	app.Listen(":3000")
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}
func dashbaordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client: %v", r.RemoteAddr)
	client := &Client{name: r.RemoteAddr, events: make(chan *DashBoard, 10)}
	go updateDashboard(client)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(1 * time.Second)
	select {
	case ev := <-client.events:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(ev)
		fmt.Fprintf(w, "data: %v\n\n", buf.String())
		fmt.Printf("data: %v\n", buf.String())
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func updateDashboard(client *Client) {
	for {
		db := &DashBoard{
			User:         uint(rand.Uint32()),
		}
		client.events <- db
	}
}
