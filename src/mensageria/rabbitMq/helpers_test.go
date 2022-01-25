package rabbitMq

import "testing"

// Teste gerado automático
func Test_gerarUrlConexao(t *testing.T) {
	type args struct {
		usuario string
		senha   string
		host    string
		port    string
		vhost   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"semVhost",
			args{"usuario",
				"senha",
				"host",
				"porta",
				"",
			},
			"amqp://usuario:senha@host:porta/",
		},
		{"comVhost",
			args{"usuario",
				"senha",
				"host",
				"porta",
				"vhost",
			},
			"amqp://usuario:senha@host:porta/vhost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gerarUrlConexao(tt.args.usuario, tt.args.senha, tt.args.host, tt.args.port, tt.args.vhost); got != tt.want {
				t.Errorf("gerarUrlConexao() = %v, want %v", got, tt.want)
			}
		})
	}
}
