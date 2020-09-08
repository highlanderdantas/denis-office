package log

import "log"

//Printa uma mensagem
func LogPrinter(message string) {
	log.Println("[denis-office]", message)
}

//Printa dois espaços
func LogSpace(spaces ...int) {
	var len = len(spaces)
	for i := 0; i <= len; i++ {
		LogPrinter("")
	}
}
