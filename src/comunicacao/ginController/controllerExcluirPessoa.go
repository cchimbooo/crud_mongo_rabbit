package ginController

import (
	"github.com/gin-gonic/gin"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
)

func excluirPessoa(p corePessoas.CorePessoaInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		if id == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "id não informado"})
			return
		}
		if errExcluir := p.Excluir(id); errExcluir != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(errExcluir))
			return
		}
		c.Status(204)
		return
	}
}
