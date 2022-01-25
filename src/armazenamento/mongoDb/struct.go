package mongoDb

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testentopus/src/core/corePessoas"
)

type pessoa struct {
	Id     IdString `bson:"_id,omitempty"`
	Nome   string   `bson:"nome,omitempty"`
	Sexo   uint8    `bson:"sexo,omitempty"`
	Peso   float64  `bson:"peso,omitempty"`
	Altura float64  `bson:"altura,omitempty"`
	IMC    float64  `bson:"imc,omitempty"`
}

// Converte da struct do mongo db para a do coreDb
func pessoaMongoParaCore(mPessoa *pessoa) *corePessoas.Pessoa {
	return &corePessoas.Pessoa{
		Id:     string(mPessoa.Id),
		Nome:   mPessoa.Nome,
		Sexo:   mPessoa.Sexo,
		Peso:   mPessoa.Peso,
		Altura: mPessoa.Altura,
		IMC:    mPessoa.IMC,
	}
}

// Converte da pessoaCore para a pessoa do Mongo
func pessoaCoreParaMongo(cPessoa *corePessoas.Pessoa) (*pessoa, error) {
	if cPessoa == nil {
		return nil, errors.New("deve-se informar uma pessoa e não nil")
	}
	p := pessoa{
		Nome:   cPessoa.Nome,
		Sexo:   cPessoa.Sexo,
		Peso:   cPessoa.Peso,
		Altura: cPessoa.Altura,
		IMC:    cPessoa.IMC,
	}

	if cPessoa.Id != "" {
		p.Id = IdString(cPessoa.Id)

	}
	return &p, nil
}

// Implementa Marshal no object ID para facilitar a vida na hora de voltar os dados
// Não faço ideia se isso seja uma má prática em mongo.
type IdString string

func (id IdString) MarshalBSONValue() (bsontype.Type, []byte, error) {
	p, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return bsontype.Null, nil, err
	}
	return bson.MarshalValue(p)
}
