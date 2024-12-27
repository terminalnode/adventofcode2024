package web

import (
	"encoding/json"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"net/http"
)

func addPrefix(prefix string, url string) string {
	if prefix == "" {
		return url
	}
	return fmt.Sprintf("/%s%s", prefix, url)
}

func readInput(r *http.Request) (util.AocInput, error) {
	var input util.AocInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return util.AocInput{}, err
	}
	defer r.Body.Close()
	return input, nil
}
