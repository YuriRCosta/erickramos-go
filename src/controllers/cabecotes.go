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

// CriarCabecote cria um cabecote no banco de dados
func CriarCabecote(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var cabecote models.Cabecote
	if err = json.Unmarshal(corpoRequisicao, &cabecote); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeCabecotes(db)
	cabecote.ID, err = repositorio.Criar(cabecote)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, cabecote)
}

// BuscarCabecotes busca todos os cabecotes salvos no banco de dados
func BuscarCabecotes(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeCabecotes(db)
	cabecotes, err := repositorio.BuscarCabecotes()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, cabecotes)
}

// BuscarCabecotePorID busca um cabecote salvo no banco de dados pelo seu ID
func BuscarCabecotePorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	cabecoteID, err := strconv.ParseUint(parametros["cabecoteID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeCabecotes(db)
	cabecote, err := repositorio.BuscarCabecotePorID(cabecoteID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, cabecote)
}

// AtualizarCabecote atualiza um cabecote no banco de dados
func AtualizarCabecote(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	cabecoteID, err := strconv.ParseUint(parametros["cabecoteID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var cabecote models.Cabecote
	if err = json.Unmarshal(corpoRequisicao, &cabecote); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeCabecotes(db)
	if err = repositorio.Atualizar(cabecoteID, cabecote); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarCabecote deleta um cabecote do banco de dados
func DeletarCabecote(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	cabecoteID, err := strconv.ParseUint(parametros["cabecoteID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeCabecotes(db)
	if err = repositorio.Deletar(cabecoteID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarCabecotePorNome busca cabecotes salvos no banco de dados que contenham o nome informado
func BuscarCabecotePorNome(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nomeCabecote := parametros["nomeCabecote"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeCabecotes(db)
	cabecote, err := repositorio.BuscarCabecotePorNome(nomeCabecote)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, cabecote)
}
