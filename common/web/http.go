package web

import (
	"errors"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"log"
	"net/http"
	"os"
)

func CreateHttpServer(
	day int,
	part1 util.Solution,
	part2 util.Solution,
) *http.Server {
	prefix := os.Getenv("AOC2024_PREFIX")

	server := &http.Server{Addr: ":8080", Handler: nil}
	addHealthCheckHandlers(prefix)
	addSolutionHandlers(prefix, day, part1, part2)
	addUnknownPathHandlers()

	go func() {
		log.Printf("Starting Day #%d service on port 8080", day)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Fatal server error: %v", err)
		}
	}()
	return server
}

func addSolutionHandlers(
	prefix string,
	day int,
	part1 util.Solution,
	part2 util.Solution,
) {
	http.HandleFunc(addPrefix(prefix, "/1"), createSolutionHandler(day, 1, part1))
	http.HandleFunc(addPrefix(prefix, "/2"), createSolutionHandler(day, 2, part2))
}

func addHealthCheckHandlers(prefix string) {
	http.HandleFunc(addPrefix(prefix, "/health"), healthCheckHandler)
	http.HandleFunc(addPrefix(prefix, "/health/live"), healthCheckHandler)
	http.HandleFunc(addPrefix(prefix, "/health/ready"), healthCheckHandler)

	if prefix != "" {
		// Add non-prefixed handlers as well
		http.HandleFunc("/health", healthCheckHandler)
		http.HandleFunc("/health/live", healthCheckHandler)
		http.HandleFunc("/health/ready", healthCheckHandler)
	}
}

func addUnknownPathHandlers() {
	http.HandleFunc("/", unknownPathHandler)
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
