package cerrors

// Database의 CRUD중 Create를 할 수 없는 경우 이 에러 구조체를 사용
type NotCreateError struct {
	ErrorMessage string
}

func (n NotCreateError) Error() string {
	return n.ErrorMessage
}
