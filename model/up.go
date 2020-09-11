package model

import (
	"fmt"
	"time"

	"github.com/highlanderdantas/denis-office/log"
)

//Up Representa a opção de subir
type Up struct {
	Operation
	Timeout int
	Amount  int
}

//UpTo Levanta os clientes gradualmente
func UpTo(up Up) {

	deploys := GetDeploysByDbNameIsDown(up.DbName)

	amount := up.Amount

	for _, deploy := range deploys {
		scaleUp(deploy.Namespace, deploy.Name)

		amount--
		if amount == 0 {
			amount = up.Amount
			setTimeout(up.Timeout)
		}
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
