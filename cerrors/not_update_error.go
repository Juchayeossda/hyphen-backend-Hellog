package cerrors

type NotUpdateError struct {
	ErrorMessage string
}

func (n NotUpdateError) Error() string {
	return n.ErrorMessage
}
