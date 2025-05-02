package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"
)

func main() {
	// Configure Keep-Alive settings
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		IdleTimeout:    60 * time.Second, // Keeps connections open for x seconds
		MaxHeaderBytes: 1 << 20,          // 1 MB max header size
	}

	fmt.Printf("[INFO][main] server running on port: %s\n", "8080")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("[ERROR][main] ListenAndServe failed: %v\n", err)
		}
	}()

	// Graceful shutdown with a timeout
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // Listen for SIGINT, SIGTERM

	// Wait for interrupt signal
	<-quit // Block until a signal is received
	handleGracefulShutdown(server)
}

func handleGracefulShutdown(server *http.Server) {
	fmt.Println("[INFO][main] Received shutdown signal")
	// Attempt to gracefully shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// stops accepting new requests, waits for inflight requests
	fmt.Printf("[INFO][main] Stops accepting new requests while waiting for in-flight requests to finish processing.\n")
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("[ERROR][main] Server forced to shut down: %v\n", err)
	}
	// Countdown before final exit
	for i := 174000; i > 0; i-- {
		fmt.Printf("[INFO][main] Exiting in %d...\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("[INFO][main] Server exiting 👋")
}
