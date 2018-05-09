package goProbability

type probError struct {
	s string
}

func (e probError) Error() string {
	return e.s
}

var (
	InvalidInput   = probError{"Invalid input"}
	NumberTooLarge = probError{"Number too large "}
)
