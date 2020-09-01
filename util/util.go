package util

import "strings"

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func ConvertToArray(value []byte) []string {
	var clientOut = string(value)
	var clients = strings.Split(clientOut, "\n")
	clients = removeIndex(clients, 0)

	return clients
}
