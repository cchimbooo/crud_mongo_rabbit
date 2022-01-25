package mongoDb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testentopus/src/core/corePessoas"
	"time"
)

func (m mongoDb) Atualizar(pessoa *corePessoas.Pessoa) (erro error) {
	if pessoa == nil {
		return errors.New("não foi informada nenhuma pessoa para ser atualizada")
	}

	if pessoa.Id == "" {
		return errors.New("a pessoa a ser atualizada não foi devidamente identificada")
	}

	// converte a pessoa do core para a pessoa da interface

	dbPessoa, pessoaErr := pessoaCoreParaMongo(pessoa)
	if pessoaErr != nil {
		return pessoaErr
	}

	// converte o id de string para objectId
	objectId, errObject := primitive.ObjectIDFromHex(pessoa.Id)
	if errObject != nil {
		return errObject
	}

	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()

	resultado, errUpdate := m.col.UpdateByID(ctx, objectId, bson.M{"$set": dbPessoa})
	if errUpdate != nil {
		return errUpdate
	}
	if resultado.MatchedCount == 0 {
		return errors.New("nenhum usuário existe com o id informado")
	}

	return nil
}
