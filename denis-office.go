package main

import (
	"fmt"
	"strings"

	"./log"
	"./model"
	"./util"
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
			pauseClients()
		case "3":
			upClients()
		case "4":
			finish()
		}
	}
}

//Le a opção escolhida
func scannerOptions() string {

	log.LogPrinter("Escolha uma das opções abaixo: ")
	log.LogSpace()
	log.LogPrinter("1 - Listar clientes")
	log.LogPrinter("2 - Pausar clientes")
	log.LogPrinter("3 - Subir clientes")
	log.LogPrinter("4 - Sair")

	var opcao string
	fmt.Scanln(&opcao)

	return opcao
}

//Lista todos clientes do cluster
func listClients() {

	clients, err := model.GetClients()

	util.ErrorHandler(err)

	clientsModel := model.ConvertTo(clients)
	for _, client := range clientsModel {
		log.LogPrinter(client.ToString())
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()
}

//Pausa todos clientes ou todos os clientes baseados em um banco
func pauseClients() {
	pause := model.Pause{}
	log.LogPrinter("Deseja parar todos os clientes: (S/N)")
	fmt.Scan(&pause.All)

	if strings.EqualFold(pause.All, "N") {
		log.LogPrinter("Deseja parar todos de qual banco.")
		log.LogPrinter("1 - dboci2")
		log.LogPrinter("2 - dboci3")
		log.LogPrinter("3 - dbocit1")
		fmt.Scan(&pause.DbName)

		pause.SetDbName()
		model.PauseTo(pause.DbName)
	} else {
		model.PauseAll()
	}

}

//Levanta todos clientes baseado num banco e um tempo
func upClients() {
	up := model.Up{}
	log.LogPrinter("Deseja subir quantos clientes de forma gradual: (3/5)")
	fmt.Scan(&up.Amount)

	log.LogPrinter("De quanto em quantos minutos: (3/5)")
	fmt.Scan(&up.Timeout)

	log.LogPrinter("Deseja parar todos de qual banco.")
	log.LogPrinter("1 - dboci2")
	log.LogPrinter("2 - dboci3")
	log.LogPrinter("3 - dbocit1")
	fmt.Scan(&up.DbName)

	up.SetDbName()

	model.UpTo(up)
}

func finish() {
	log.LogPrinter("Encerrando programa!!")
}
