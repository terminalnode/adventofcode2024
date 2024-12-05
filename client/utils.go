package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func readFile(path string) string {
	data, err := readFileSafely(path)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return data
}

func readFileSafely(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file '%s' with error: %v", path, err)
	}

	buff := bytes.NewBuffer(data)
	content, err := io.ReadAll(buff)
	if err != nil {
		return "", fmt.Errorf("failed to parse data from file '%s' with error: %v", path, err)
	}

	return string(content), nil
}

func saveFileSafely(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0744)
	return err
}

func fetchInput(
	day int,
	sessionToken string,
) (string, error) {
	cachePath := os.Getenv("AOC2024_INPUT_CACHE_PATH")
	var cacheFile string
	cacheEnabled := false
	if cachePath != "" {
		cacheEnabled = true
		cacheFile = fmt.Sprintf("%s/day%d-%s", cachePath, day, sessionToken)

		content, err := readFileSafely(cacheFile)
		if err == nil {
			return content, nil
		}
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	strBody, _ := strings.CutSuffix(string(body), "\n")

	if cacheEnabled {
		_ = saveFileSafely(cacheFile, strBody)
	}

	return strBody, nil
}
