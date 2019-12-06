package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

// TODO: Melhorar para utilizar variáveis de ambiente nas configurações
var configs = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "root",
	Database: "goapi",
}

// Sess que é uma variável que faz a conexão com o banco de dados
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(configs)

	if err != nil {
		log.Fatal(err.Error())
	}
}
