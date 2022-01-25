package mongoDb

import (
	"errors"
	"regexp"
	"strings"
	"testentopus/src/utils/erroSimples"
)

func gerarUrlMongo(usuario, senha, host, porta, opcoesDeConexao string) string {
	b := strings.Builder{}
	b.WriteString("mongodb://")
	if usuario != "" {
		b.WriteString(usuario)
		b.WriteString(":")
		b.WriteString(senha)
		b.WriteString("@")
	}
	b.WriteString(host)
	b.WriteString(":")
	b.WriteString(porta)
	if opcoesDeConexao != "" {
		b.WriteString("/")
		b.WriteString(opcoesDeConexao)
	}

	return b.String()
}

func limparObjectID(s string) (string, error) {
	re := regexp.MustCompile(`"(.*?)"`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		return "", erroSimples.GerarErro(errors.New("não foi possível obter o id da pessoa"), 500)
	}
	return match[1], nil

}
