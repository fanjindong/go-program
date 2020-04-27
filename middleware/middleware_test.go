package middleware

type Middleware interface {
	BeforeCall(...interface{}) error
	AfterCall(...interface{}) error
}

type Log struct {
}

func (l *Log)

type Service struct {
	mws []
}

func (s *Service) Sum(x, y int) int {
	return x + y
}

func (s *Service) Multiply(x, y int) int {
	return x * y
}

