package mongoDb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testentopus/src/core/corePessoas"
	"time"
)

// Se fosse usar para mais coisas provavelmetne quebraria isso em conexões menores cada uma com uma interface
// Teria uma factory pai que teria métodos para retornar as interfaces do main.

func MongoDbFactory(db, collection, usuario, senha, host, porta, opcoesDeConexao string) (corePessoas.ArmazenamentoInterface, error) {
	// Cria url
	url := gerarUrlMongo(usuario, senha, host, porta, opcoesDeConexao)
	fmt.Println(url)
	// Conmecta ao mongo
	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()
	client, errConn := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if errConn != nil {
		return nil, errConn
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	colecao := client.Database(db).Collection(collection)
	return mongoDb{cliente: client, col: colecao}, nil
}

type mongoDb struct {
	cliente *mongo.Client
	col     *mongo.Collection
}

func (m mongoDb) Desconectar() {
	if err := m.cliente.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
