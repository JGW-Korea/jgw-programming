package math

type Calculator interface {
	Add(a int, b int) int
	subtract(a int, b int) int
}

type Math struct{}

func (math Math) Add(a int, b int) int {
	return a + b
}

func (math Math) subtract(a int, b int) int {
	return a - b
}