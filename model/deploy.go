package model

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/highlanderdantas/denis-office/log"
	"github.com/highlanderdantas/denis-office/util"
)

//Deploy representa uma deployment no kubernetes
type Deploy struct {
	Namespace string
	Name      string
}

//GetDeploysBy pega todos os deploys baseado num banco online ou offline
func GetDeploysBy(dbName string, isDown bool) []Deploy {
	var grep string

	if isDown {
		grep = "| grep 0/0"
	} else {
		grep = "| grep -e 1/1 -e 0/1"
	}

	command := fmt.Sprintf("kubectl get deploy --no-headers -A -l cloud.dbName=%s %s| awk '{print $1 \",\" $2}' ", dbName, grep)
	cmd, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)

	return ConvertToDeploy(cmd)
}

//GetDeploysByDbNameIsDown pega todos os deploys de um banco offline
func GetDeploysByDbNameIsDown(dbName string) []Deploy {
	return GetDeploysBy(dbName, true)
}

//GetDeploysByDbNameIsUp pega todos os deploys de um banco online
func GetDeploysByDbNameIsUp(dbName string) []Deploy {
	return GetDeploysBy(dbName, false)
}

//GetDeploy pega todos os deploys
func GetDeploy() []Deploy {

	command := fmt.Sprintf("kubectl get deploy --no-headers -A | grep -e 1/1 -e 0/1 | awk '{print $1 \",\" $2}' ")
	cmd, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)

	return ConvertToDeploy(cmd)
}

//ConvertToDeploy converte um array de bytes em um array de Deploy
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

//isEmpty verifica se a variavel
func isEmpty(value []string) bool {
	return value != nil && value[0] != "" && value[1] != ""
}

//scaleDown abaixa a scala de um deployment de um determinado namespace
func scaleDown(namespace string, name string) {
	command := getScaleBy(name, namespace, false)

	message := fmt.Sprintf("Pausando %s do cliente %s", name, namespace)
	log.LogPrinter(message)

	_, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)
}

//scaleUp abaixa a scala de um deployment de um determinado namespace
func scaleUp(namespace string, name string) {
	command := getScaleBy(name, namespace, true)

	message := fmt.Sprintf("Subindo %s do cliente %s", name, namespace)
	log.LogPrinter(message)

	_, err := exec.Command("bash", "-c", command).Output()

	util.ErrorHandler(err)
}

//getScaleBy monta o comando para abaixar ou subir escala de um deployment
func getScaleBy(name string, namespace string, isUp bool) string {
	var count int
	if isUp {
		count = 1
	}

	return fmt.Sprintf("kubectl scale deploy/%s -n %s  --replicas %d ", name, namespace, count)
}
