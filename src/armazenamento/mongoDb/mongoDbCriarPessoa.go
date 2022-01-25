package mongoDb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testentopus/src/core/corePessoas"
	"time"
)

func (m mongoDb) Criar(cPessoa *corePessoas.Pessoa) (string, error) {
	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()

	// Converte da entidade de pessoa para a entidade de usada pelo mongo
	pes, errPessoa := pessoaCoreParaMongo(cPessoa)
	if errPessoa != nil {
		return "", errPessoa
	}
	// Chama método de inserir
	id, err := m.col.InsertOne(ctx, pes)
	if err != nil {
		return "", err
	}

	// Converte o Id do mongo para string para o resto do sistema ler
	s := id.InsertedID.(primitive.ObjectID).String()
	// ObjectID("61ef21ebdd2a99e30dad2ef7")
	stringLimpa, errLimpar := limparObjectID(s)
	if errLimpar != nil {
		return "", errLimpar
	}
	return stringLimpa, nil
}
