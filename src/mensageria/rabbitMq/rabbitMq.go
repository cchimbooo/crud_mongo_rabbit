package rabbitMq

import (
	"github.com/streadway/amqp"
	"testentopus/src/core/corePessoas"
)

func RabbitMqFactory(nomeFila, usuario, senha, host, port, vhost string) (corePessoas.MensageriaInterface, error) {
	// Conecta com o rabbit
	conn, errDial := amqp.Dial(gerarUrlConexao(usuario, senha, host, port, vhost))
	if errDial != nil {
		return nil, errDial
	}

	// Abre conexão com um channel
	canal, errCha := conn.Channel()
	if errCha != nil {
		return nil, errCha
	}

	fila, errFila := canal.QueueDeclare(
		nomeFila,
		true,
		false,
		false,
		false,
		nil,
	)
	if errFila != nil {
		return nil, errFila
	}
	return rabbit{ch: canal, qu: fila}, nil
}

type rabbit struct {
	ch *amqp.Channel
	qu amqp.Queue
}

func (r rabbit) Publicar(mensagem []byte) error {
	return r.ch.Publish(
		"",
		r.qu.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        mensagem,
		})
}
