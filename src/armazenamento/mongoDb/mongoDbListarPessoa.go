package mongoDb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testentopus/src/core/corePessoas"
	"time"
)

func (m mongoDb) Listar(_ corePessoas.FiltroPessoas) ([]corePessoas.Pessoa, error) {

	var listaDbPessoa []pessoa

	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()

	// Faz a busca
	retorno, errFind := m.col.Find(ctx, bson.M{})
	// Da o decode para o model de listaDbPessoa
	if errFind != nil {
		return nil, errFind
	}
	if err := retorno.All(ctx, &listaDbPessoa); err != nil {
		log.Fatal(err)
	}

	// cria slice de pessoaCore
	l := len(listaDbPessoa)
	listaCorePessoa := make([]corePessoas.Pessoa, l, l)
	for k, _ := range listaDbPessoa {
		listaCorePessoa[k] = *pessoaMongoParaCore(&listaDbPessoa[k])
	}

	//// jeito mais rápido mas que futuramente alguém vai fazer quebrar sem saber.
	// listaCorePessoa := *(*[]pessoa)(unsafe.Pointer(&listaDbPessoa))

	return listaCorePessoa, nil
}
