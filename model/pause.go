package model

import (
	"fmt"

	"../log"
)

//Representa a opção de pausar
type Pause struct {
	All    string
	DbName string
}

//Pausa todos os deploys online com uma tag
func PauseTo(dbName string) {

	deploys := GetDeploysByDbNameIsUp(dbName)

	for _, deploy := range deploys {
		scaleDown(deploy.Namespace, deploy.Name)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()

}

func PauseAll() {

	deploys := GetDeploy()

	for _, deploy := range deploys {
		scaleDown(deploy.Namespace, deploy.Name)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()
}
