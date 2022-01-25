package corePessoas

import (
	"bytes"
	"errors"
	"testentopus/src/utils/erroSimples"
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
		return erroSimples.GerarErro(errors.New("nome da pessoa deve ser preenchido"), 422)
	}

	if p.Sexo == 0 {
		return erroSimples.GerarErro(errors.New("sexo da pessoa deve ser preenchido"), 422)
	}

	if p.Peso == 0 {
		return erroSimples.GerarErro(errors.New("peso da pessoa deve ser preenchido"), 422)
	}
	if p.Altura == 0 {
		return erroSimples.GerarErro(errors.New("altura da pessoa deve ser preenchida"), 422)
	}

	if p.IMC == 0 {
		return erroSimples.GerarErro(errors.New("imc deve ser preenchido"), 422)
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
