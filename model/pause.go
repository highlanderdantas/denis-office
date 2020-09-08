package model

import (
	"fmt"
	"os/exec"
	"strings"

	"../log"
	"../util"
)

//Representa a opção de pausar
type Pause struct {
	All    string
	DbName string
}

//Representa uma deployment no kubernetes
type Deploy struct {
	Namespace string
	Name      string
}

//Pausa todos os deploys online com uma tag
func PauseTo(dbName string) {

	deploys := getDeploysBy(dbName)

	for _, deploy := range deploys {
		scaleDown(deploy.Namespace, deploy.Name)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()

}

func PauseAll() {

	deploys := getDeploy()

	for _, deploy := range deploys {
		scaleDown(deploy.Namespace, deploy.Name)
	}

	log.LogSpace(2)
	log.LogPrinter("Digite enter para voltar para o menu")
	fmt.Scanln()
}

//Abaixa a scala de um deployment de um determinado namespace
func scaleDown(namespace string, name string) {
	command := fmt.Sprintf("kubectl scale deploy/%s -n %s  --replicas 0 ", name, namespace)

	message := fmt.Sprintf("Pausando %s do cliente %s", name, namespace)
	log.LogPrinter(message)

	_, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)

	message = fmt.Sprintf("%s do cliente %s pausado", name, namespace)
	log.LogPrinter(message)

}

//Pega todos os deploys baseado num banco
func getDeploysBy(dbName string) []Deploy {

	command := fmt.Sprintf("kubectl get deploy --no-headers -A -l cloud.dbName=%s | grep -e 1/1 -e 0/1 | awk '{print $1 \",\" $2}' ", dbName)
	cmd, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)

	return ConvertToDeploy(cmd)
}

//Pega todos os deploys
func getDeploy() []Deploy {

	command := fmt.Sprintf("kubectl get deploy --no-headers -A | grep -e 1/1 -e 0/1 | awk '{print $1 \",\" $2}' ")
	cmd, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)

	return ConvertToDeploy(cmd)
}

//Converte um array de bytes em um array de Deploy
func ConvertToDeploy(value []byte) []Deploy {
	deploysOut := string(value)
	deploysStd := strings.Split(deploysOut, "\n")

	var deploys []Deploy
	for _, deploy := range deploysStd {
		dp := strings.Split(deploy, ",")
		if isEmpty(dp) {
			deploys = append(deploys, Deploy{
				Namespace: dp[0],
				Name:      strings.TrimSpace(dp[1]),
			})
		}
	}

	return deploys
}

func isEmpty(dp []string) bool {
	return dp != nil && dp[0] != "" && dp[1] != ""
}
