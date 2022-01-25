package mongoDb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
	"time"
)

func (m mongoDb) Atualizar(pessoa *corePessoas.Pessoa) (erro error) {
	if pessoa == nil {
		return erroSimples.GerarErro(errors.New("não foi informada nenhuma pessoa para ser atualizada"), 422)
	}

	if pessoa.Id == "" {
		return erroSimples.GerarErro(errors.New("a pessoa a ser atualizada não foi devidamente identificada"), 422)
	}

	// converte a pessoa do core para a pessoa da interface

	dbPessoa, pessoaErr := pessoaCoreParaMongo(pessoa)
	if pessoaErr != nil {
		return erroSimples.GerarErro(pessoaErr, 500, "falha ler dados da pessoa")
	}

	// converte o id de string para objectId
	objectId, errObject := primitive.ObjectIDFromHex(pessoa.Id)
	if errObject != nil {
		return erroSimples.GerarErro(errObject, 422, "impossível ler o Id da pessoa")
	}

	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()

	resultado, errUpdate := m.col.UpdateByID(ctx, objectId, bson.M{"$set": dbPessoa})
	if errUpdate != nil {
		return erroSimples.GerarErro(errUpdate, 500, "falha ao atualizar a pessoa")
	}
	if resultado.MatchedCount == 0 {
		return erroSimples.GerarErro(errors.New("nenhum usuário existe com o id informado"), 400)
	}

	return nil
}
