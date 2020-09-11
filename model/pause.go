package model

import (
	"fmt"

	"github.com/highlanderdantas/denis-office/log"
)

//Pause Representa a opção de pausar
type Pause struct {
	Operation
	All string
}

//PauseTo Pausa todos os deploys online com uma tag
func PauseTo(dbName string) {

	deploys := GetDeploysByDbNameIsUp(dbName)

	for _, deploy := range deploys {
		scaleDown(deploy.Namespace, deploy.Name)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()

}

//PauseAll pausa todos os clientes
func PauseAll() {

	deploys := GetDeploy()

	for _, deploy := range deploys {
		scaleDown(deploy.Namespace, deploy.Name)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()
}
