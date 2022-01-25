package corePessoas

import (
	"bytes"
	"errors"
)

type Pessoa struct {
	Id     string
	Nome   string
	Sexo   uint8 // 1 -Feminino 2 - Masculino  3 - Não Identificado ... definir
	Peso   float64
	Altura float64
	IMC    float64
}

func (p Pessoa) validarPessoa() error {
	if p.Nome == "" {
		return errors.New("nome da pessoa deve ser preenchido")
	}

	if p.Sexo == 0 {
		return errors.New("sexo da pessoa deve ser preenchido")
	}

	if p.Peso == 0 {
		return errors.New("peso da pessoa deve ser preenchido")
	}
	if p.Altura == 0 {
		return errors.New("altura da pessoa deve ser preenchida")
	}

	if p.IMC == 0 {
		return errors.New("imc deve ser preenchido")
	}
	return nil
}

func (p Pessoa) dadosParaPublicar() []byte {
	return []byte(p.Id)
}

func (p Pessoa) mensagemCriar() []byte {
	b := bytes.Buffer{}
	b.WriteString("Cadastro de pessoa ")
	b.Write(p.dadosParaPublicar())
	return b.Bytes()
}

func (p Pessoa) mensagemAtualizar() []byte {
	b := bytes.Buffer{}
	b.WriteString("Edição de pessoa ")
	b.Write(p.dadosParaPublicar())
	return b.Bytes()
}

type FiltroPessoas struct {
}
