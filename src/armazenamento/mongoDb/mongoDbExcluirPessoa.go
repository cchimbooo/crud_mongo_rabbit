package mongoDb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testentopus/src/utils/erroSimples"
	"time"
)

func (m mongoDb) Excluir(id string) error {

	// converte o id de string para objectId
	objectId, errObject := primitive.ObjectIDFromHex(id)
	if errObject != nil {
		return errObject
	}

	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()

	result, errDel := m.col.DeleteOne(ctx, bson.M{"_id": objectId})
	if errDel != nil {
		return erroSimples.GerarErro(errDel, 422, "falha ao deletar pessoa")
	}
	if result.DeletedCount == 0 {
		return erroSimples.GerarErro(errors.New("o id informado não existia no banco"), 422)
	}
	return nil
}
