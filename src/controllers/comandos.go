package controllers

import (
	"encoding/json"
	"erickramos-go/src/banco"
	"erickramos-go/src/models"
	"erickramos-go/src/repositories"
	"erickramos-go/src/respostas"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarComando cria um comando no banco de dados
func CriarComando(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var comando models.Comando
	if err = json.Unmarshal(corpoRequisicao, &comando); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeComandos(db)
	comando.ID, err = repositorio.Criar(comando)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, comando)
}

// BuscarComandos busca todos os comandos salvos no banco de dados
func BuscarComandos(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeComandos(db)
	comandos, err := repositorio.BuscarComandos()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, comandos)
}

// BuscarComandoPorID busca um comando salvo no banco de dados pelo seu ID
func BuscarComandoPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	comandoID, err := strconv.ParseUint(parametros["comandoID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeComandos(db)
	comando, err := repositorio.BuscarComandoPorID(comandoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, comando)
}

// AtualizarComando atualiza um comando no banco de dados
func AtualizarComando(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	comandoID, err := strconv.ParseUint(parametros["comandoID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var comando models.Comando
	if err = json.Unmarshal(corpoRequisicao, &comando); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeComandos(db)
	if err = repositorio.Atualizar(comandoID, comando); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarComando deleta um comando do banco de dados
func DeletarComando(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	comandoID, err := strconv.ParseUint(parametros["comandoID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeComandos(db)
	if err = repositorio.Deletar(comandoID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarComandoPorNome busca comandos salvos no banco de dados que contenham o nome informado
func BuscarComandoPorNome(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nomeComando := parametros["nomeComando"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeComandos(db)
	comando, err := repositorio.BuscarComandoPorNome(nomeComando)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, comando)
}
