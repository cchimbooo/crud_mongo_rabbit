package corePessoas

import (
	"bytes"
	"errors"
)

// PessoaCoreFactory Instancia uma CorePessoaInterface
func PessoaCoreFactory(db ArmazenamentoInterface, mensageria MensageriaInterface) (CorePessoaInterface, error) {
	if db == nil {
		return nil, errors.New("interface de armazenamento não encontrada")
	}
	if mensageria == nil {
		return nil, errors.New("interface de mensageria não encontrada")
	}
	return pessoaCore{db: db, qu: mensageria}, nil
}

// pessoaCore implementa uma CorePessoaInterface
type pessoaCore struct {
	db ArmazenamentoInterface
	qu MensageriaInterface
}

func (p pessoaCore) Criar(pessoa *Pessoa) (string, error) {
	// valida se pessoa esta completa
	if pessoa == nil {
		return "", errors.New("pessoa não informada")
	}

	// Cria a pessoa e pega o id  -- TODO ver se tem que retornar o ID ou o objeto
	id, errCriar := p.db.Criar(pessoa)
	if errCriar != nil {
		return "", errCriar
	}
	// Garante que a pessoa tenha o mesmo ID para caso de uma implementação de db ser diferente.
	// e o db não insira o código no model.
	pessoa.Id = id
	// Publica e retorna eventual erro se ocorrer.
	errPub := p.qu.Publicar(pessoa.mensagemCriar())
	return id, errPub
}

func (p pessoaCore) Atualizar(pessoa *Pessoa) error {
	if pessoa == nil {
		return errors.New("pessoa não informada")
	}
	if pessoa.Id == "" {
		return errors.New("id da pessoa não informado")
	}
	// Valida se esta preenchida
	if err := pessoa.validarPessoa(); err != nil {
		return err
	}
	// Atualiza a pessoa
	if err := p.db.Atualizar(pessoa); err != nil {
		return err
	}
	// Publica a atualização
	if err := p.qu.Publicar(pessoa.mensagemAtualizar()); err != nil {
		return err
	}
	return nil
}

func (p pessoaCore) Excluir(id string) error {

	if id == "" {
		return errors.New("deve ser informado um ID para a exclusão")
	}
	// Exclui a pessoa
	if err := p.db.Excluir(id); err != nil {
		return err
	}
	// builer da mensagem, esta aqui pq não tem pessoa
	b := bytes.Buffer{}
	b.WriteString("Exclusão de pessoa ")
	b.WriteString(id)
	// Publica a mensagem
	if err := p.qu.Publicar(b.Bytes()); err != nil {
		return err
	}
	return nil
}

func (p pessoaCore) Obter(id string) (*Pessoa, error) {
	if id == "" {
		return nil, errors.New("id deve ser preenchido")
	}
	return p.db.Obter(id)
}

func (p pessoaCore) Listar(filtroPessoas FiltroPessoas) ([]Pessoa, error) {
	return p.db.Listar(filtroPessoas)
}
