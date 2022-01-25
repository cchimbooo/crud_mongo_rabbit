package corePessoas

import (
	"errors"
	"regexp"
)

type ArmazenamentoInterfaceMock struct {
}

func (d ArmazenamentoInterfaceMock) Criar(pessoa *Pessoa) (id string, erro error) {
	if pessoa == nil || pessoa.Peso == 1 {
		return "", errors.New("erroMock")
	}
	if pessoa.Nome == "errado" {
		return "errado", nil
	}
	return "certo", nil
}

func (d ArmazenamentoInterfaceMock) Atualizar(pessoa *Pessoa) (erro error) {
	if pessoa == nil || pessoa.Peso == 1 {
		return errors.New("erroMock")
	}
	return nil
}

func (d ArmazenamentoInterfaceMock) Excluir(id string) error {
	if id == "1" {
		return errors.New("erroMock")
	}
	return nil
}

func (d ArmazenamentoInterfaceMock) Obter(id string) (*Pessoa, error) {
	if id == "1" {
		return nil, errors.New("erroMock")
	}
	if id == "errado" {
		return &Pessoa{Id: "errado"}, nil
	}

	return &Pessoa{Id: "certo"}, nil
}

func (d ArmazenamentoInterfaceMock) Listar(_ FiltroPessoas) ([]Pessoa, error) {
	return []Pessoa{}, nil
}

func (d ArmazenamentoInterfaceMock) Desconectar() {
}

type MensageriaInterfaceMock struct {
}

func (m MensageriaInterfaceMock) Publicar(bytes []byte) error {

	if match, _ := regexp.Match(`.*errado.*`, bytes); match {
		return errors.New("erro Mock")
	}
	return nil
}
