package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

//cookie
var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config.HaskKey, config.BlockKey)
}

// registar as informacoes de autenticacao
func Salvar(w http.ResponseWriter, ID string, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/", //para funcionar em tudo
		HttpOnly: true,
	})

	return nil
}
