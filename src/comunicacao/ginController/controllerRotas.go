package ginController

import (
	"github.com/gin-gonic/gin"
	"log"
	"testentopus/src/core/corePessoas"
)

func InicializarRotas(corePessoa corePessoas.CorePessoaInterface) {
	if corePessoa == nil {
		log.Fatalln("Falha ao inicializar o core de pessoa")
	}
	r := gin.Default()
	pessoa := r.Group("/pessoa")

	pessoa.POST("", criarPessoa(corePessoa))
	pessoa.PUT("/:id", atualizarPessoa(corePessoa))
	pessoa.GET("", listarPessoa(corePessoa))
	pessoa.GET("/:id", obterPessoa(corePessoa))
	pessoa.DELETE("/:id", excluirPessoa(corePessoa))

	if err := r.Run(); err != nil {
		log.Fatalln("Falha ao rodar o servidor")
	}
}
