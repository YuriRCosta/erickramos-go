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

// CriarJunta cria um junta no banco de dados
func CriarJunta(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var junta models.Junta
	if err = json.Unmarshal(corpoRequisicao, &junta); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeJuntas(db)
	junta.ID, err = repositorio.Criar(junta)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, junta)
}

// BuscarJuntas busca todos os juntas salvos no banco de dados
func BuscarJuntas(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeJuntas(db)
	juntas, err := repositorio.BuscarJuntas()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, juntas)
}

// BuscarJuntaPorID busca um junta salvo no banco de dados pelo seu ID
func BuscarJuntaPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	juntaID, err := strconv.ParseUint(parametros["juntaID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeJuntas(db)
	junta, err := repositorio.BuscarJuntaPorID(juntaID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, junta)
}

// AtualizarJunta atualiza um junta no banco de dados
func AtualizarJunta(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	juntaID, err := strconv.ParseUint(parametros["juntaID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var junta models.Junta
	if err = json.Unmarshal(corpoRequisicao, &junta); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeJuntas(db)
	if err = repositorio.Atualizar(juntaID, junta); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarJunta deleta um junta do banco de dados
func DeletarJunta(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	juntaID, err := strconv.ParseUint(parametros["juntaID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeJuntas(db)
	if err = repositorio.Deletar(juntaID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarJuntaPorCabecotes busca juntas salvos no banco de dados que contenham o cabecotes informado
func BuscarJuntaPorCabecotes(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	cabecotesJunta := parametros["cabecotesJunta"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeJuntas(db)
	junta, err := repositorio.BuscarJuntaPorCabecotes(cabecotesJunta)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, junta)
}
