package lib

import (
	"log"
	"os"
	_ "github.com/joho/godotenv/autoload"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

// TODO: Melhorar para utilizar variáveis de ambiente nas configurações
var configs = mysql.ConnectionURL{
	Host:     os.Getenv("DB_HOST"),
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	Database: os.Getenv("DB_DATABASE"),
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
