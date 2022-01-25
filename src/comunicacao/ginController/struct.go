package ginController

import (
	"errors"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
)

type PessoaRequest struct {
	Nome   string  `json:"nome"`
	Sexo   uint8   `json:"sexo"`
	Peso   float64 `json:"peso"`
	Altura float64 `json:"altura"`
	IMC    float64 `json:"imc"`
}

func (p PessoaRequest) validar() error {
	if p.Nome == "" {
		return erroSimples.GerarErro(errors.New("nome da pessoa em branco"), 400)
	}
	if p.Sexo == 0 {
		return erroSimples.GerarErro(errors.New("sexo da pessoa em branco"), 400)
	}

	if p.Altura == 0 {
		return erroSimples.GerarErro(errors.New("altura da pessoa em branco"), 400)
	}
	if p.IMC == 0 {
		return erroSimples.GerarErro(errors.New("imc da pessoa em branco"), 400)
	}
	return nil
}

func (p PessoaRequest) converterParaCore(id *string) *corePessoas.Pessoa {
	cp := corePessoas.Pessoa{
		Nome:   p.Nome,
		Sexo:   p.Sexo,
		Peso:   p.Peso,
		Altura: p.Altura,
		IMC:    p.IMC,
	}
	if id != nil {
		cp.Id = *id
	}
	return &cp
}

type PessoaRessource struct {
	Id     string  `json:"id"`
	Nome   string  `json:"nome"`
	Sexo   uint8   `json:"sexo"`
	Peso   float64 `json:"peso"`
	Altura float64 `json:"altura"`
	IMC    float64 `json:"imc"`
}

func corePessoaParaHttpPessoaRessource(p *corePessoas.Pessoa) (PessoaRessource, error) {
	if p == nil {
		return PessoaRessource{}, erroSimples.GerarErro(errors.New("não foi possível obter a pessoa"), 500)
	}
	return PessoaRessource{
		Id:     p.Id,
		Nome:   p.Nome,
		Sexo:   p.Sexo,
		Peso:   p.Peso,
		Altura: p.Altura,
		IMC:    p.IMC,
	}, nil
}

func (p PessoaRessource) converterParaCore() corePessoas.Pessoa {
	return corePessoas.Pessoa{
		Id:     p.Id,
		Nome:   p.Nome,
		Sexo:   p.Sexo,
		Peso:   p.Peso,
		Altura: p.Altura,
		IMC:    p.IMC,
	}
}
