package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Solution = func(string) string

func Setup(
	day int,
	part1 Solution,
	part2 Solution,
) {
	http.HandleFunc("/1", createSolutionHandler(day, 1, part1))
	http.HandleFunc("/2", createSolutionHandler(day, 2, part2))
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/health/live", healthCheckHandler)
	http.HandleFunc("/health/ready", healthCheckHandler)

	fmt.Printf("Starting Day #%d service on port 3000\n", day)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)

	}
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
