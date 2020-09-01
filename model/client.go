package model

import (
	"fmt"

	"github.com/highlanderdantas/denis-office/util"
)

type Client struct {
	Name string
}

func ConvertToArray(stdout []byte) []Client {

	values := util.ConvertToArray(stdout)
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

func (c Client) ToString() string {
	return fmt.Sprint("name: ", c.Name)
}
