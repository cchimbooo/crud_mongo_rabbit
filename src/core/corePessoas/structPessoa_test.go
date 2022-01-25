package corePessoas

import (
	"testing"
)

func TestPessoa_validarPessoa(t *testing.T) {
	type fields struct {
		Id     string
		Nome   string
		Sexo   uint8
		Peso   float64
		Altura float64
		IMC    float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"SemNome",
			fields{
				Id:     "1",
				Nome:   "",
				Sexo:   1,
				Peso:   80.5,
				Altura: 1.75,
				IMC:    221,
			},
			true,
		},
		{
			"SemSexo",
			fields{
				Id:     "1",
				Nome:   "FElipe",
				Sexo:   0,
				Peso:   80.5,
				Altura: 1.75,
				IMC:    221,
			},
			true,
		},
		{
			"SemPeso",
			fields{
				Id:     "1",
				Nome:   "FElipe",
				Sexo:   2,
				Peso:   0,
				Altura: 1.75,
				IMC:    221,
			},
			true,
		}, {
			"SemAltura",
			fields{
				Id:     "1",
				Nome:   "FElipe",
				Sexo:   2,
				Peso:   80.5,
				Altura: 0,
				IMC:    221,
			},
			true,
		}, {
			"SemImc",
			fields{
				Id:     "1",
				Nome:   "FElipe",
				Sexo:   3,
				Peso:   80.5,
				Altura: 1.75,
				IMC:    0,
			},
			true,
		}, {
			"Completo",
			fields{
				Id:     "1",
				Nome:   "FElipe",
				Sexo:   2,
				Peso:   80.5,
				Altura: 1.75,
				IMC:    221,
			},
			false,
		}, {
			"limpo",
			fields{
				Id:     "",
				Nome:   "",
				Sexo:   0,
				Peso:   0,
				Altura: 0,
				IMC:    0,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pessoa{
				Id:     tt.fields.Id,
				Nome:   tt.fields.Nome,
				Sexo:   tt.fields.Sexo,
				Peso:   tt.fields.Peso,
				Altura: tt.fields.Altura,
				IMC:    tt.fields.IMC,
			}
			if err := p.validarPessoa(); (err != nil) != tt.wantErr {
				t.Errorf("validarPessoa() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
