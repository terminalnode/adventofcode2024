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
	fmt.Println("Args", args)

	flagSet := flag.NewFlagSet("run", flag.ExitOnError)
	day := flagSet.Int("day", 0, "The day to solve")
	part := flagSet.Int("part", 0, "The part of the provided day to solve")
	file := flagSet.String("file", "", "Path to the file containing the input data.")

	err := flagSet.Parse(args)
	if err != nil {
		fmt.Printf("Failed to parse flags: %v\n", err)
		os.Exit(1)
	}

	validateDayAndPart(*day, *part)

	var input string
	switch {
	case *file != "":
		input = readFile(*file)
	default:
		fmt.Printf("No input method selected, can't get puzzle input.\n")
		os.Exit(1)
	}

	answer, err := runSolution(*day, *part, input)
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
	day int,
	part int,
	input string,
) (string, error) {
	fmt.Printf("Running day %d, part %d\n", day, part)
	url := fmt.Sprintf("http://localhost:%d/%d", 3000+day, part)
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

	if resp.StatusCode != 200 && bodyReadErr == nil {
		return "", fmt.Errorf("unexpected status code %d", resp.StatusCode)
	} else if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code %d: '%s'", resp.StatusCode, body)
	}

	return string(body), nil
}
