package router

import "github.com/gorilla/mux"

//Gerar retorna um router com os pacotes configurados
func Gerar() *mux.Router {
	return mux.NewRouter()
}
