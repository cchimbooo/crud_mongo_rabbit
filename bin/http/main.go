package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testentopus/src/armazenamento/mongoDb"
	"testentopus/src/comunicacao/ginController"
	"testentopus/src/core/corePessoas"
	"testentopus/src/mensageria/rabbitMq"
	"time"
)

func main() {
	if loadError := godotenv.Load(); loadError != nil {
		log.Fatalln(loadError.Error())
	}
	time.Sleep(30 * time.Second)
	armazenamento, errArmazenamento := mongoDb.MongoDbFactory(
		os.Getenv("DBBANCO"),
		os.Getenv("DBCOLLECTION"),
		os.Getenv("DBUSUARIO"),
		os.Getenv("DBSENHA"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORTA"),
		os.Getenv("DBOPCOES"),
	)
	if errArmazenamento != nil {
		log.Fatalln(errArmazenamento.Error())
	}

	mensageria, errMensageria := rabbitMq.RabbitMqFactory(
		os.Getenv("RABBITFILA"),
		os.Getenv("RABBITUSER"),
		os.Getenv("RABBITPWD"),
		os.Getenv("RABBITHOST"),
		os.Getenv("RABBITPORT"),
		os.Getenv("RABBITVHOST"),
	)
	if errMensageria != nil {
		log.Fatalln(errMensageria.Error())
	}

	core, errCore := corePessoas.PessoaCoreFactory(armazenamento, mensageria)
	if errCore != nil {
		log.Fatalln(errCore.Error())
	}
	ginController.InicializarRotas(core)
}
