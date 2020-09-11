package model

import "../log"

//Representa uma operação
type Operation struct {
	DbName string
}

//Seta um banco
func (o *Operation) SetDbName() {
	switch o.DbName {
	case "1":
		log.LogPrinter("teste")
		o.DbName = "dboci2"
	case "2":
		o.DbName = "dboci3"
	case "3":
		o.DbName = "dbocit1"
	}
}
