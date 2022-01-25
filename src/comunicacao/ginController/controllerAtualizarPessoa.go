package ginController

import (
	"github.com/gin-gonic/gin"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
)

func atualizarPessoa(p corePessoas.CorePessoaInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		var pessoa PessoaRequest
		id := c.Params.ByName("id")

		if id == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "id não informado"})
			return
		}

		if err := c.ShouldBindJSON(&pessoa); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "impossível ler o body do request"})
			return
		}

		if err := pessoa.validar(); err != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(err))
			return
		}
		if errAtualizar := p.Atualizar(pessoa.converterParaCore(&id)); errAtualizar != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(errAtualizar))
			return
		}
		c.Status(204)
		return
	}
}
