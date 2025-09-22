package logging

import (
	"fmt"
)

/*func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}*/

func LogInfo(msg string) {
	fmt.Println("Log : ", msg)
}

func LogError(msg string) {
	fmt.Println("Error : ", msg)
}
