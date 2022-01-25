package mongoDb

import (
	"errors"
	"fmt"
	"testentopus/src/core/corePessoas"
	"testing"
)

// Rotina de teste, pressupõem que o banco esteja vazio e possua uma uniqueindex em pessoa
func TestRotina(t *testing.T) {
	var banco = "teste"
	var collection = "teste"
	var usuario = ""
	var senha = ""
	var host = "localhost"
	var porta = "27017"
	var opcoesDeConexao = ""

	db, errDb := MongoDbFactory(banco, collection, usuario, senha, host, porta, opcoesDeConexao)
	if errDb != nil {
		t.Errorf("falha ao se coneta ao banco")
	}
	t.Run("Teste rotina banco", func(t *testing.T) {
		if err := rotina(db); err != nil {
			t.Errorf(err.Error())
		}

	})
}

// Por se tratar de um teste de integração não fiz usando o padrão de testes do go.
// trata-se de uma rotina que testa as funcionalidades da interface
func rotina(db corePessoas.ArmazenamentoInterface) error {
	// FEcha o banco
	defer db.Desconectar()

	// Cria modelo
	felipe := &corePessoas.Pessoa{
		Nome:   "Felipe",
		Sexo:   2,
		Peso:   76.56,
		Altura: 1.70,
		IMC:    123,
	}

	lista, errListaBase := db.Listar(corePessoas.FiltroPessoas{})
	if errListaBase != nil {
		return errListaBase
	}
	if len(lista) != 0 {
		return errors.New("teste deve ser rodado com banco limpo")
	}

	// Cria pessoa 1
	id, errCriar1 := db.Criar(felipe)
	if errCriar1 != nil {
		return errCriar1
	}

	// Atribui o id para pessoa pois o codigo não faz isso.
	felipe.Id = id

	// Pega retorno do banco
	felipeRetorno, errObeter := db.Obter(id)
	if errObeter != nil {
		return errObeter
	}

	// Compara se objeto de retonro é == ao objeto inserido + ID
	if !eIgual(felipe, felipeRetorno) {
		return errors.New("documento inserido não é == ao documento obtido")
	}

	// Lista os objetos
	lista, errLista := db.Listar(corePessoas.FiltroPessoas{})
	if errLista != nil {
		return errLista
	}

	// Verifica se tem algo na lista
	if len(lista) == 0 {
		return errors.New("lista vazia")
	}

	// Verifica se o objeto 1 é == o valor inserido.
	if !eIgual(&lista[0], felipe) {
		return errors.New("lista na posição 0 não é a esperada")
	}

	// Altera o nome e atualiza
	felipe.Nome = "Nome atualizado"

	if errAtt := db.Atualizar(felipe); errAtt != nil {
		return errAtt
	}

	// Consulta no banco
	obter2, errObter2 := db.Obter(id)
	if errObter2 != nil {
		return errObter2
	}

	// verifica se atualizou
	if !eIgual(felipe, obter2) {
		return errors.New("atualizar não deu certo, o obter é diferente do valor passado")
	}

	// Limpa o ID para fazer um registro novo que deve ser rejeitado por ter o mesmo nome
	felipe.Id = ""
	_, errInserir2 := db.Criar(felipe)
	if errInserir2 == nil {
		return errors.New("não gerou erro ao criar um nome com o mesmo nome")
	}

	// exclui o registro
	if errExc := db.Excluir(id); errExc != nil {
		return errExc
	}

	// Cria uma nova pessoa com o mesmo nome para verificar se excluiu
	felipe.Id = ""
	id3, errInserir3 := db.Criar(felipe)
	if errInserir3 != nil {
		return errInserir3
	}
	// Cria uma pessoa com o nome novo para ter dois registros
	felipe.Nome = "novo"
	id4, errInserir4 := db.Criar(felipe)
	if errInserir4 != nil {
		return errInserir4
	}

	// Pega o id da 2ª pessoa do banco e tenta atualizar o nome para a primeira para ver se bloqueia
	felipe.Id = id4
	felipe.Nome = "Nome atualizado"
	if errUpdate := db.Atualizar(felipe); errUpdate == nil {
		return errors.New("erro ao atualizar com o mesmo nome")
	}

	// Limpa o banco para poder rodar o teste de novo
	if err := db.Excluir(id3); err != nil {
		return err
	}
	if err := db.Excluir(id4); err != nil {
		return err
	}
	return nil
}

func eIgual(a, b *corePessoas.Pessoa) bool {
	if a.Id != b.Id {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	if a.Nome != b.Nome {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	if a.Sexo != b.Sexo {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	if a.Peso != b.Peso {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	if a.Altura != b.Altura {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	if a.IMC != b.IMC {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	return true
}
