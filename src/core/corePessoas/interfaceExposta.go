package corePessoas

type CorePessoaInterface interface {
	Criar(pessoa *Pessoa) (string, error)
	Atualizar(pessoa *Pessoa) error
	Excluir(id string) error
	Obter(id string) (*Pessoa, error)
	Listar(FiltroPessoas) ([]Pessoa, error)
}
