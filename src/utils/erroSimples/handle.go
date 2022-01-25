package erroSimples

import "errors"

func Handle(err error) (int, map[string]interface{}) {

	if err == nil {
		return 204, nil
	}
	var e ErroSimplesInterface
	if !errors.As(err, &e) {
		e = GerarErro(err, 500)
	}
	return e.Codigo(), e.Retorno()

}
