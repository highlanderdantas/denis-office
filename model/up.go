package model

import (
	"fmt"
	"time"

	"../log"
)

//Representa a opção de subir
type Up struct {
	Operation
	Timeout int
	Amount  int
}

//Levanta os clientes gradualmente
func UpTo(up Up) {

	deploys := GetDeploysByDbNameIsDown(up.DbName)

	for _, deploy := range deploys {
		scaleUp(deploy.Namespace, deploy.Name)
		setTimeout(up.Timeout)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()
}

//Seta um intervalo de tempo
func setTimeout(timeout int) {
	duration := time.Duration(timeout) * time.Minute
	time.Sleep(duration)
}
