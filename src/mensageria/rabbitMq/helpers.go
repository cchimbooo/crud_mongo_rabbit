package rabbitMq

import "strings"

func gerarUrlConexao(usuario, senha, host, port, vhost string) string {
	b := strings.Builder{}
	b.WriteString("amqp://")
	b.WriteString(usuario)
	b.WriteString(":")
	b.WriteString(senha)
	b.WriteString("@")
	b.WriteString(host)
	b.WriteString(":")
	b.WriteString(port)
	b.WriteString("/")
	b.WriteString(vhost)
	return b.String()
}
