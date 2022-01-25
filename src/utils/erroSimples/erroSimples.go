package erroSimples

type ErroSimplesInterface interface {
	error
	Retorno() map[string]interface{}
	Codigo() int
}

func GerarErro(err error, codigo int, mensagem ...string) ErroSimplesInterface {
	if err == nil {
		return nil
	}
	var msg string
	if len(mensagem) > 0 {
		msg = mensagem[0]
	} else {
		msg = err.Error()
	}
	e := erroStruct{
		erroBase: err,
		mensagem: msg,
		codigo:   codigo,
	}
	return e
}

type erroStruct struct {
	erroBase error
	mensagem string
	codigo   int
}

func (e erroStruct) Error() string {
	return e.erroBase.Error()
}

func (e erroStruct) Retorno() map[string]interface{} {
	if e.codigo == 404 {
		return nil
	}
	return map[string]interface{}{"error": e.mensagem}
}

func (e erroStruct) Codigo() int {
	return e.codigo
}
