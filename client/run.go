package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func runCmd(args []string) {
	flagSet := flag.NewFlagSet("run", flag.ExitOnError)
	day := flagSet.Int("day", 0, "The day to solve")
	part := flagSet.Int("part", 0, "The part of the provided day to solve")
	file := flagSet.String("file", "", "Path to the file containing the input data.")
	url := flagSet.String("url", "localhost", "The URL from which to call the backend.")

	err := flagSet.Parse(args)
	if err != nil {
		fmt.Printf("Failed to parse flags: %v\n", err)
		os.Exit(1)
	}

	validateDayAndPart(*day, *part)

	sessionToken := os.Getenv("AOC2024_SESSION_TOKEN")

	var input string
	switch {
	case *file != "":
		input = readFile(*file)
	case sessionToken != "":
		data, err := fetchInput(*day, sessionToken)
		if err != nil {
			fmt.Printf("Failed to fetch input using session token '%s': %v", sessionToken, err)
			os.Exit(1)
		}
		input = data
	default:
		fmt.Println("No input method selected and session token not set, can't get puzzle input.")
		os.Exit(1)
	}

	answer, err := runSolution(*url, *day, *part, input)
	if err != nil {
		fmt.Printf("Failed to run solution: %v", err)
		os.Exit(1)
	}
	fmt.Println(answer)
}

func validateDayAndPart(day int, part int) {
	if day < 1 || day > 25 {
		fmt.Printf("The day has to be a number between 1 and 25, but was '%d'\n", day)
		os.Exit(1)
	} else if part < 1 || part > 2 {
		fmt.Printf("The part has to be a number between 1 and 2, but was '%d'\n", part)
		os.Exit(1)
	}
}

func runSolution(
	url string,
	day int,
	part int,
	input string,
) (string, error) {
	url = fmt.Sprintf("http://%s/day%02d/%d", url, day, part)
	fmt.Printf("Running day %d, part %d with URL '%s'\n", day, part, url)
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(input)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, bodyReadErr := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 && bodyReadErr != nil {
		return "", fmt.Errorf("unexpected status code %d", resp.StatusCode)
	} else if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code %d: '%s'", resp.StatusCode, body)
	}

	return string(body), nil
}
