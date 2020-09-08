package model

import (
	"fmt"
	"os/exec"

	"../util"
)

//Representa um cliente
type Client struct {
	Name string
}

//Converte um array de byte em um array de Clientes
func ConvertTo(stdout []byte) []Client {

	values := util.ConvertToByte(stdout)
	clients := []Client{}

	for _, value := range values {
		if value != "" {
			clients = append(clients, Client{
				Name: value,
			})
		}
	}

	return clients
}

//Printa o nome do cliente
func (c Client) ToString() string {
	return fmt.Sprint("name: ", c.Name)
}

//Pega todos clientes
func GetClients() ([]byte, error) {
	cmd := exec.Command("bash", "-c", "kubectl get namespaces --no-headers -l cattle.io/creator=norman | cut -d \" \" -f 1 ")
	return cmd.Output()
}
