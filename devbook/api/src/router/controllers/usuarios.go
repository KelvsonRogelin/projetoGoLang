package controllers

import "net/http"

//CriarUsuario insere usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}

//BuscarUsuarios busca todos usuarios no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar todos Usuarios"))
}

//BuscarUsuario busca um usuario no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar um Usuario"))
}

//AtualizarUsuario atualiza dados de um usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario"))
}

//DeletarUsuario Deleta dados de um usuario no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Usuario"))
}
