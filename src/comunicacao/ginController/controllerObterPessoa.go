package ginController

import (
	"github.com/gin-gonic/gin"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
)

func obterPessoa(p corePessoas.CorePessoaInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		if id == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "id não informado"})
			return
		}
		pessoa, errObter := p.Obter(id)
		if errObter != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(errObter))
			return
		}

		pRessource, errConverter := corePessoaParaHttpPessoaRessource(pessoa)
		if errConverter != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(errConverter))
			return
		}

		c.JSON(200, gin.H{"pessoa": pRessource})
		return
	}
}
