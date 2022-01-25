package ginController

import (
	"github.com/gin-gonic/gin"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
)

func criarPessoa(p corePessoas.CorePessoaInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		var pessoa PessoaRequest
		if err := c.ShouldBindJSON(&pessoa); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "impossível ler o body do request"})
			return
		}
		if err := pessoa.validar(); err != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(err))
			return
		}
		id, errCriar := p.Criar(pessoa.converterParaCore(nil))
		if errCriar != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(errCriar))
			return
		}
		c.JSON(200, gin.H{"id": id})
		return
	}
}
