package cerrors

type NotDeleteError struct {
	ErrorMessage string
}

func (e NotDeleteError) Error() string {
	return e.ErrorMessage
}
