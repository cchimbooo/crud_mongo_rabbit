package mongoDb

import "testing"

func Test_gerarUrlMongo(t *testing.T) {
	type args struct {
		usuario         string
		senha           string
		host            string
		porta           string
		opcoesDeConexao string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"teste1",
			args{
				usuario:         "a",
				senha:           "a",
				host:            "a",
				porta:           "a",
				opcoesDeConexao: "a",
			},
			"mongodb://a:a@a:a/a",
		},
		{"teste2",
			args{
				usuario:         "",
				senha:           "",
				host:            "a",
				porta:           "a",
				opcoesDeConexao: "a",
			},
			"mongodb://a:a/a",
		},
		{"teste3",
			args{
				usuario:         "a",
				senha:           "a",
				host:            "a",
				porta:           "a",
				opcoesDeConexao: "",
			},
			"mongodb://a:a@a:a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gerarUrlMongo(tt.args.usuario, tt.args.senha, tt.args.host, tt.args.porta, tt.args.opcoesDeConexao); got != tt.want {
				t.Errorf("gerarUrlMongo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limparObjectID(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"teste", args{s: `teste"asd"`}, "asd", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := limparObjectID(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("limparObjectID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("limparObjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
