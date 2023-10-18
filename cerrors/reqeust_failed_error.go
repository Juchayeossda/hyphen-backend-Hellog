package cerrors

type ReqeustFailedError struct {
	ErrorMessage string
}

func (r ReqeustFailedError) Error() string {
	return r.ErrorMessage
}
