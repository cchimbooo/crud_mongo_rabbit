package rabbitMq

import (
	"fmt"
	"strings"
)

func gerarUrlConexao(usuario, senha, host, port, vhost string) string {
	b := strings.Builder{}
	b.WriteString("amqp://")
	b.WriteString(usuario)
	b.WriteString(":")
	b.WriteString(senha)
	b.WriteString("@")
	b.WriteString(host)
	if port != "" {
		b.WriteString(":")
	}
	b.WriteString(port)
	b.WriteString("/")
	b.WriteString(vhost)
	fmt.Println(b.String())
	return b.String()
}
