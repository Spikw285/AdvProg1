package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	mu         sync.Mutex
	data       map[string]string
	requests   int
	shutdownCh chan struct{}
}

// NewServer creates and initializes a new Server
func NewServer() *Server {
	return &Server{
		data:       make(map[string]string),
		shutdownCh: make(chan struct{}),
	}
}

// POST /data: Adds data to the in-memory database
func (s *Server) postDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	for key, value := range body {
		s.data[key] = value
	}
	s.requests++
	s.mu.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Data stored successfully")
}

// GET /data: Returns the entire database as JSON
func (s *Server) getDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	s.mu.Lock()
	response, _ := json.Marshal(s.data)
	s.requests++
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GET /stats: Returns the number of handled requests
func (s *Server) statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	s.mu.Lock()
	stats := map[string]int{"requests": s.requests}
	s.requests++
	s.mu.Unlock()

	response, _ := json.Marshal(stats)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// DELETE /data/{key}: Deletes a specific key from the database
func (s *Server) deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Path[len("/data/"):]

	s.mu.Lock()
	if _, exists := s.data[key]; exists {
		delete(s.data, key)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Deleted key: %s\n", key)
	} else {
		http.Error(w, "Key not found", http.StatusNotFound)
	}
	s.requests++
	s.mu.Unlock()
}

// Background worker logs server status every 5 seconds
func (s *Server) startBackgroundWorker() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			log.Printf("Server status: %d requests, %d data entries", s.requests, len(s.data))
			s.mu.Unlock()
		case <-s.shutdownCh:
			ticker.Stop()
			log.Println("Background worker stopped")
			return
		}
	}
}

// Graceful shutdown of the server
func (s *Server) shutdown() {
	close(s.shutdownCh)
	log.Println("Server is shutting down gracefully...")
}

func main() {
	server := NewServer()

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			server.postDataHandler(w, r)
		case http.MethodGet:
			server.getDataHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/data/", server.deleteDataHandler)
	http.HandleFunc("/stats", server.statsHandler)

	// Start background worker
	go server.startBackgroundWorker()

	// Start the server
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Set up channel to listen for system signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-signalChan
	log.Printf("Received signal: %s, shutting down...", sig)
	server.shutdown()
}
