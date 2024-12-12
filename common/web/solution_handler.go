package web

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"net/http"
)

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

func defaultSolutionHandler(
	day int,
	part int,
) util.Solution {
	return func(input string) string {
		return fmt.Sprintf("Solution for day %d part %d not implemented yet", day, part)
	}
}
