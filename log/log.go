package log

import "log"

func LogPrinter(message string) {
	log.Println("[denis-office]", message)
}

func LogSpace(spaces ...int) {
	var len = len(spaces)
	for i := 0; i <= len; i++ {
		LogPrinter("")
	}
}
