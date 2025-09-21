package service

import (
	"fmt"
)

type GreeterService struct {
	tp TimeProvider
}

func NewGreeterService(tp TimeProvider) *GreeterService {
	return &GreeterService{tp: tp}
}

func (gs *GreeterService) Greet(name string) string {
	h := gs.tp.Now().Hour()
	switch {
	case h < 12:
		return fmt.Sprintf("Good morning %s", name)
	case h < 18:
		return fmt.Sprintf("Good afternoon %s", name)

	default:
		return fmt.Sprintf("Good night %s", name)

	}
}

/*type GreeterService struct {
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}
func (gs *GreeterService) Greet(name string) string {
	t := time.Now().Hour()
	switch {
	case t < 12:
		return fmt.Sprintf("Good morning %s", name)

	case t < 18:
		return fmt.Sprintf("Good afternoon %s", name)

	default:
		return fmt.Sprintf("Good night %s", name)
	}
}
*/
