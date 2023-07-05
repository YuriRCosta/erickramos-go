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

// CriarValvula cria um valvula no banco de dados
func CriarValvula(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var valvula models.Valvula
	if err = json.Unmarshal(corpoRequisicao, &valvula); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	valvula.ID, err = repositorio.Criar(valvula)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, valvula)
}

// BuscarValvulas busca todos os valvulas salvos no banco de dados
func BuscarValvulas(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	valvulas, err := repositorio.BuscarValvulas()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, valvulas)
}

// BuscarValvulaPorID busca um valvula salvo no banco de dados pelo seu ID
func BuscarValvulaPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	valvulaID, err := strconv.ParseUint(parametros["valvulaID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	valvula, err := repositorio.BuscarValvulaPorID(valvulaID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, valvula)
}

// AtualizarValvula atualiza um valvula no banco de dados
func AtualizarValvula(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	valvulaID, err := strconv.ParseUint(parametros["valvulaID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var valvula models.Valvula
	if err = json.Unmarshal(corpoRequisicao, &valvula); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	if err = repositorio.Atualizar(valvulaID, valvula); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarValvula deleta um valvula do banco de dados
func DeletarValvula(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	valvulaID, err := strconv.ParseUint(parametros["valvulaID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	if err = repositorio.Deletar(valvulaID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarValvulaPorNome busca valvulas salvos no banco de dados que contenham o nome informado
func BuscarValvulaPorNome(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nomeValvula := parametros["nomeValvula"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	valvula, err := repositorio.BuscarValvulaPorNome(nomeValvula)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, valvula)
}

// BuscarValvulaPorTipo busca valvulas salvos no banco de dados que contenham a tipo informada
func BuscarValvulaPorTipo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tipoValvula := parametros["tipoValvula"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeValvulas(db)
	valvula, err := repositorio.BuscarValvulaPorTipo(tipoValvula)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, valvula)
}
