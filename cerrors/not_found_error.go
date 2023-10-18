package cerrors

type NotFoundError struct {
	ErrorMessage string
}

func (n NotFoundError) Error() string {
	return n.ErrorMessage
}
