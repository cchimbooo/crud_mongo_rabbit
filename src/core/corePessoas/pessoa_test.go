package corePessoas

import (
	"reflect"
	"testing"
)

func TestPessoaCoreFactory(t *testing.T) {
	type args struct {
		db         ArmazenamentoInterface
		mensageria MensageriaInterface
	}
	tests := []struct {
		name    string
		args    args
		want    CorePessoaInterface
		wantErr bool
	}{
		{"certo",
			args{ArmazenamentoInterfaceMock{}, MensageriaInterfaceMock{}},
			pessoaCore{ArmazenamentoInterfaceMock{}, MensageriaInterfaceMock{}},
			false,
		},
		{"err1",
			args{ArmazenamentoInterfaceMock{}, nil},
			nil,
			true,
		},
		{"err2",
			args{nil, MensageriaInterfaceMock{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PessoaCoreFactory(tt.args.db, tt.args.mensageria)
			if (err != nil) != tt.wantErr {
				t.Errorf("PessoaCoreFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PessoaCoreFactory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pessoaCore_Atualizar(t *testing.T) {
	type fields struct {
		db ArmazenamentoInterface
		qu MensageriaInterface
	}
	f := fields{ArmazenamentoInterfaceMock{}, MensageriaInterfaceMock{}}

	type args struct {
		pessoa *Pessoa
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"vazio", f, args{pessoa: nil}, true},
		{"ErroPreenchimento", f, args{pessoa: &Pessoa{}}, true},
		{"Erro Db", f, args{pessoa: &Pessoa{"1", "1", 1, 1, 1, 1}}, true},
		{"Erro Mock", f, args{pessoa: &Pessoa{"errado", "errado", 2, 2, 2, 2}}, true},
		{"certo", f, args{pessoa: &Pessoa{"certo", "certo", 2, 2, 2, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pessoaCore{
				db: tt.fields.db,
				qu: tt.fields.qu,
			}
			if err := p.Atualizar(tt.args.pessoa); (err != nil) != tt.wantErr {
				t.Errorf("Atualizar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pessoaCore_Criar(t *testing.T) {
	type fields struct {
		db ArmazenamentoInterface
		qu MensageriaInterface
	}
	f := fields{ArmazenamentoInterfaceMock{}, MensageriaInterfaceMock{}}

	type args struct {
		pessoa *Pessoa
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"ErradoPreenchimento", f, args{pessoa: &Pessoa{}}, "", true},
		{"erroNil", f, args{pessoa: nil}, "", true},
		{"erroDb", f, args{pessoa: &Pessoa{"certo", "certo", 1, 1, 1, 1}}, "", true},
		{"erroRabbit", f, args{pessoa: &Pessoa{"errado", "errado", 1, 2, 1, 1}}, "errado", true},
		{"certo", f, args{pessoa: &Pessoa{"certo", "certo", 2, 2, 2, 2}}, "certo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pessoaCore{
				db: tt.fields.db,
				qu: tt.fields.qu,
			}
			got, err := p.Criar(tt.args.pessoa)
			if (err != nil) != tt.wantErr {
				t.Errorf("Criar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Criar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pessoaCore_Excluir(t *testing.T) {
	type fields struct {
		db ArmazenamentoInterface
		qu MensageriaInterface
	}

	f := fields{ArmazenamentoInterfaceMock{}, MensageriaInterfaceMock{}}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"vazio", f, args{id: ""}, true},
		{"DbErro", f, args{id: "1"}, true},
		{"RabbitErro", f, args{id: "errado"}, true},
		{"certo", f, args{id: "certo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pessoaCore{
				db: tt.fields.db,
				qu: tt.fields.qu,
			}
			if err := p.Excluir(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Excluir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pessoaCore_Obter(t *testing.T) {
	type fields struct {
		db ArmazenamentoInterface
		qu MensageriaInterface
	}

	f := fields{ArmazenamentoInterfaceMock{}, MensageriaInterfaceMock{}}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Pessoa
		wantErr bool
	}{
		{"vazio", f, args{id: ""}, nil, true},
		{"erroDb", f, args{id: "1"}, nil, true},
		{"certo", f, args{id: "certo"}, &Pessoa{Id: "certo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pessoaCore{
				db: tt.fields.db,
				qu: tt.fields.qu,
			}
			got, err := p.Obter(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Obter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Obter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
