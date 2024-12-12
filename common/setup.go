package common

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Solution = func(string) string

func Setup(
	day int,
	part1 Solution,
	part2 Solution,
) {
	prefix := os.Getenv("AOC2024_PREFIX")
	http.HandleFunc(addPrefix(prefix, "/1"), createSolutionHandler(day, 1, part1))
	http.HandleFunc(addPrefix(prefix, "/2"), createSolutionHandler(day, 2, part2))
	http.HandleFunc(addPrefix(prefix, "/health"), healthCheckHandler)
	http.HandleFunc(addPrefix(prefix, "/health/live"), healthCheckHandler)
	http.HandleFunc(addPrefix(prefix, "/health/ready"), healthCheckHandler)

	if prefix != "" {
		// For health endpoints, add non-prefixed handlers as well
		http.HandleFunc("/health", healthCheckHandler)
		http.HandleFunc("/health/live", healthCheckHandler)
		http.HandleFunc("/health/ready", healthCheckHandler)
	}

	http.HandleFunc("/", unknownPathHandler)

	// Run the HTTP server on port 8080 in the background
	server := &http.Server{Addr: ":8080", Handler: nil}
	go func() {
		log.Printf("Starting Day #%d service on port 8080", day)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Fatal server error: %v", err)
		}
	}()

	// Open a signal channel, listening for SIGTERM and SIGINT
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	log.Printf("Received signal %s, shutting down...", <-signalChan)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func addPrefix(prefix string, url string) string {
	if prefix == "" {
		return url
	}
	return fmt.Sprintf("/%s%s", prefix, url)
}

func createSolutionHandler(
	day int,
	part int,
	solution func(string) string,
) func(http.ResponseWriter, *http.Request) {
	if solution == nil {
		solution = defaultSolutionHandler(day, part)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := whitelistMethods([]string{"POST"}, w, r); err != nil {
			fmt.Print(err.Error())
			return
		}

		input, err := readInput(r)
		if err != nil {
			http.Error(w, "Failed to read input", http.StatusBadRequest)
			return
		}

		result := solution(input)
		if _, err = w.Write([]byte(result)); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
	}
}

func healthCheckHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := whitelistMethods([]string{"GET", "POST"}, w, r); err != nil {
		fmt.Print(err.Error())
		return
	}

	if _, err := w.Write([]byte("{ \"status\": \"UP\" }")); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func unknownPathHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	msg := fmt.Sprintf("Invalid path: %s", r.URL.Path)
	http.Error(w, msg, http.StatusNotFound)
}

func defaultSolutionHandler(
	day int,
	part int,
) Solution {
	return func(input string) string {
		return fmt.Sprintf("Solution for day %d part %d not implemented yet", day, part)
	}
}

func whitelistMethods(
	methods []string,
	w http.ResponseWriter,
	r *http.Request,
) error {
	for _, method := range methods {
		if r.Method == method {
			return nil
		}
	}

	http.Error(
		w,
		fmt.Sprintf("%s is not in allowed methods: %q", r.Method, methods),
		http.StatusMethodNotAllowed,
	)
	return fmt.Errorf("expected method to be one of %q, but was %s", methods, r.Method)
}

func readInput(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	return string(body), nil
}
