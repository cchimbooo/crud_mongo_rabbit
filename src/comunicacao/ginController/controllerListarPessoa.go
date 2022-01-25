package ginController

import (
	"github.com/gin-gonic/gin"
	"testentopus/src/core/corePessoas"
	"testentopus/src/utils/erroSimples"
)

func listarPessoa(p corePessoas.CorePessoaInterface) func(c *gin.Context) {
	return func(c *gin.Context) {

		pessoasCore, errListar := p.Listar(corePessoas.FiltroPessoas{})
		if errListar != nil {
			c.AbortWithStatusJSON(erroSimples.Handle(errListar))
			return
		}
		tamanho := len(pessoasCore)
		pessoaListaRetorno := make([]PessoaRessource, tamanho, tamanho)
		for k, pessoaCore := range pessoasCore {
			pRessource, errConverter := corePessoaParaHttpPessoaRessource(&pessoaCore)
			if errConverter != nil {
				c.AbortWithStatusJSON(erroSimples.Handle(errConverter))
				return
			}
			pessoaListaRetorno[k] = pRessource
		}

		c.JSON(200, gin.H{"data": pessoaListaRetorno})
		return
	}
}
