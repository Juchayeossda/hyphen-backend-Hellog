package exception

// error를 검사하여 에러가 있으면 panic발생
func InspectErr(err error) {
	if err != nil {
		panic(err)
	}
}
