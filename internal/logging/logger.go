package logging

import "fmt"

func LogInfo(msg string) {
	fmt.Println("Log : ", msg)
}

func LogError(msg string) {
	fmt.Println("Error : ", msg)
}
