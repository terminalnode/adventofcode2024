package common

import (
	"errors"
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
	http.HandleFunc("/1", createHandler(part1))
	http.HandleFunc("/2", createHandler(part2))
	fmt.Printf("Starting Day #%d service on port 3000\n", day)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func createHandler(solution func(string) string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := methodMustBePost(w, r); err != nil {
			fmt.Print(err.Error())
			return
		}

		input, err := readInput(r)
		if err != nil {
			http.Error(w, "Failed to read input", http.StatusBadRequest)
			return
		}

		result := solution(input)
		_, err = w.Write([]byte(result))
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
	}
}

func methodMustBePost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		msg := fmt.Sprintf("Expected method to be POST, but was %s\n", r.Method)
		return errors.New(msg)
	}

	return nil
}

func readInput(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	return string(body), nil
}
