package web

import (
	"fmt"
	"io"
	"net/http"
)

func addPrefix(prefix string, url string) string {
	if prefix == "" {
		return url
	}
	return fmt.Sprintf("/%s%s", prefix, url)
}

func readInput(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	return string(body), nil
}
