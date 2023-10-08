package error

import (
	"fmt"
)

func Error(line int, message string) {
	Report(line, "", message)
}

func Report(line int, where string, message string) {
	fmt.Println("[line " + string(rune(line)) + "] Error" + where + ": " + message)
	//hadError := true
}
