package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//go:embed index.html
var staticFiles embed.FS

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	interval := flag.Duration("interval", time.Second, "interval between SSE events")
	flag.Parse()

	logger := log.New(os.Stdout, "sse: ", log.LstdFlags|log.Lmsgprefix)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("/events", sseHandler(*interval, logger))

	srv := &http.Server{
		Addr:              *addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Run the server in the background so main can wait for a shutdown signal.
	serverErr := make(chan error, 1)
	go func() {
		logger.Printf("server is running on %s", *addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	// Wait for either a fatal server error or an interrupt/terminate signal.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	select {
	case err := <-serverErr:
		logger.Fatalf("server failed: %v", err)
	case <-ctx.Done():
		logger.Println("shutting down...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatalf("graceful shutdown failed: %v", err)
	}

	logger.Println("server stopped")
}

// sseHandler returns an http.HandlerFunc that streams a timestamped event to
// the client every interval until the client disconnects.
func sseHandler(interval time.Duration, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		rc := http.NewResponseController(w)
		if err := rc.Flush(); err != nil {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}

		logger.Printf("client connected: %s", r.RemoteAddr)
		defer logger.Printf("client disconnected: %s", r.RemoteAddr)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-r.Context().Done():
				return
			case now := <-ticker.C:
				if _, err := fmt.Fprintf(w, "data: The time is %s\n\n", now.Format(time.UnixDate)); err != nil {
					return
				}
				if err := rc.Flush(); err != nil {
					return
				}
			}
		}
	}
}
