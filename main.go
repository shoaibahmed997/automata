package main

import (
	"fmt"
	"log"
	"net/http"

	hook "github.com/robotn/gohook"
	"github.com/shoaibahmed997/automata/handler"
)

func startServer() {

	go func() {
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Server error:", err)
		}
	}()
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

func main() {
	startServer()
	handler.MainHandler()
	low()

	// Your main application logic goes here
	fmt.Println("Main application is running. Server is available at http://localhost:8080")

	// Keep the main goroutine running

	select {}
}
