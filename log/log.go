package log

import "log"

//LogPrinter printa uma mensagem
func LogPrinter(message string) {
	log.Println("[denis-office]", message)
}

//LogSpace printa dois espaços
func LogSpace(spaces ...int) {
	var len = len(spaces)
	for i := 0; i <= len; i++ {
		LogPrinter("")
	}
}
