package util

type ErrorType int

const (
	NotImplemented ErrorType = iota
	InputParsingError
	ParsingError
	StringToNumber
	ProcessingError
)

func (et ErrorType) String() string {
	switch et {
	case NotImplemented:
		return "NotImplemented"
	case ParsingError:
		return "ParsingError"
	case StringToNumber:
		return "StringToNumber"
	default:
		return "UnknownError"
	}
}

type AocError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	IsError bool   `json:"-"`
}

func (
	e AocError,
) Error() string {
	return e.Message
}
