# BUild
FROM golang:1.16.3-alpine

# Definino o direitório de trabalho, a partir de agora todos os comandos serão executados a partir desse path
WORKDIR $GOPATH/cmd/testentopus

# Copiando os fontes do projeto para a pasta cmd do GOPATH/ciot
COPY . .

# Baixando todas as dependencias do go, compilando e instalando o projeto
RUN go install testentopus/bin/http

# Abrindo a porta 8080 padrão
EXPOSE 8080

# Inializando o conteiner
CMD ../../bin/http

