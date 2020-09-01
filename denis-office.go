package main

import (
	"fmt"
	"os/exec"

	"./log"
	"./model"
)

func main() {
	var option string
	log.LogPrinter("Olá Seja bem vindo")

	for option != "4" {

		option = scannerOptions()

		switch option {
		case "1":
			listClients()
		case "2":
			upClients()
		case "3":
			pausedClients()
		case "4":
			finish()
		}
	}
}

func scannerOptions() string {

	log.LogPrinter("Escolha uma das opções abaixo: ")
	log.LogSpace()
	log.LogPrinter("1 - Listar clientes")
	log.LogPrinter("2 - Subir clientes")
	log.LogPrinter("3 - Pausar clientes")
	log.LogPrinter("4 - Sair")

	var opcao string
	fmt.Scanln(&opcao)

	return opcao
}

func listClients() {

	clients, err := exec.Command("bash", "-c", "kubectl get namespaces -l cattle.io/creator=norman | cut -d \" \" -f 1 ").Output()

	if err != nil {
		log.LogPrinter(err.Error())
		return
	}

	clientsModel := model.ConvertToArray(clients)
	for _, client := range clientsModel {
		log.LogPrinter(client.ToString())
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()
}

func upClients() {
	//implementar
}

func pausedClients() {
	//implementar
}

func finish() {
	log.LogPrinter("Encerrando programa!!")
}
