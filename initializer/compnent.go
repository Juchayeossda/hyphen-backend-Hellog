package initializer

import (
	"hyphen-backend-hellog/cerrors/exception"
	"os"

	"github.com/joho/godotenv"
)

type Component interface {
	Get(string) string
}

type IComponent struct{}

func (IComponent) Get(key string) string {
	return os.Getenv(key)
}

func NewCompnent(filenames ...string) Component {
	err := godotenv.Load(filenames...)
	exception.Sniff(err)
	return IComponent{}
}
