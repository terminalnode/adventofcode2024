package util

import "fmt"

type AocInput struct {
	Input string `json:"input"`
}

type AocSolution struct {
	Solution string `json:"solution"`
}

type Solution = func(AocInput) (AocSolution, AocError)

func NewAocError(
	m string,
	t ErrorType,
) (AocSolution, AocError) {
	return AocSolution{}, AocError{Message: m, Type: t.String()}
}

func NewAocSolution(
	solution string,
) (AocSolution, AocError) {
	return AocSolution{Solution: solution}, AocError{}
}

func FormatAocSolution(
	format string,
	a ...any,
) (AocSolution, AocError) {
	return NewAocSolution(fmt.Sprintf(format, a...))
}
