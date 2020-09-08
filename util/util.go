package util

import (
	"os"
	"strings"

	"../log"
)

//Remove um index especifico de um array
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

//Converte um array de bytes em um array de string
func ConvertToByte(value []byte) []string {
	var clientOut = string(value)
	var clients = strings.Split(clientOut, "\n")
	clients = RemoveIndex(clients, 0)

	return clients
}

//Verifica se tem um erro printa e finaliza o programa
func ErrorHandler(err error) {
	if err != nil {
		log.LogPrinter(err.Error())
		os.Exit(1)
	}
}
