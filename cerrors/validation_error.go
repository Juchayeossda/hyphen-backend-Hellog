package cerrors

type ValidationError struct {
	ErrorMessage string
}

func (v ValidationError) Error() string {
	return v.ErrorMessage
}
