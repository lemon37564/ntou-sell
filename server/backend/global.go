package backend

// TimeLayout is a format that rules what a string of time should be
const TimeLayout = "2006-01-02"

type beError struct {
	text string
}

func newBeErr(text string) beError {
	return beError{text: text}
}

func (b beError) Error() string {
	return b.text
}
