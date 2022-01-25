package corePessoas

type ArmazenamentoInterface interface {
	Criar(pessoa *Pessoa) (id string, erro error)
	Atualizar(*Pessoa) (erro error)
	Excluir(id string) error
	Obter(id string) (*Pessoa, error)
	Listar(FiltroPessoas) ([]Pessoa, error)
	Desconectar()
}

type MensageriaInterface interface {
	Publicar([]byte) error
}
