package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"net/http"
)

func createSolutionHandler(
	day int,
	part int,
	solution util.Solution,
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

		aocSolution, aocErr := solution(input)
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if errors.Is(err, util.AocError{}) {
			encoder.Encode(aocErr)
		} else {
			encoder.Encode(aocSolution)
		}
	}
}

func defaultSolutionHandler(
	day int,
	part int,
) util.Solution {
	return func(input util.AocInput) (util.AocSolution, util.AocError) {
		return util.NewAocError(
			fmt.Sprintf("Solution for day %d part %d not implemented yet", day, part),
			util.NotImplemented,
		)
	}
}
