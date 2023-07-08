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

// CriarRetentorComando cria um retentor de comando no banco de dados
func CriarRetentorComando(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var retentorComando models.RetentorComando
	if err = json.Unmarshal(corpoRequisicao, &retentorComando); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresComandoRepository(db)
	retentorComando.ID, err = repositorio.Criar(retentorComando)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, retentorComando)
}

// BuscarRetentoresComando busca todos os retentores de comando salvos no banco de dados
func BuscarRetentoresComando(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresComandoRepository(db)
	retentoresComando, err := repositorio.BuscarRetentoresComando()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	respostas.JSON(w, http.StatusOK, retentoresComando)
}

// BuscarRetentorComandoPorID busca um retentor de comando salvo no banco de dados pelo seu ID
func BuscarRetentorComandoPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorComandoID, err := strconv.ParseUint(parametros["retentorComandoID"], 10, 64)
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

	repositorio := repositories.NewRetentoresComandoRepository(db)
	retentorComando, err := repositorio.BuscarRetentorComandoPorID(retentorComandoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, retentorComando)
}

// AtualizarRetentorComando atualiza um retentor de comando no banco de dados
func AtualizarRetentorComando(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorComandoID, err := strconv.ParseUint(parametros["retentorComandoID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var retentorComando models.RetentorComando
	if err = json.Unmarshal(corpoRequisicao, &retentorComando); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresComandoRepository(db)
	if err = repositorio.Atualizar(retentorComandoID, retentorComando); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarRetentorComando deleta um retentor de comando do banco de dados
func DeletarRetentorComando(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorComandoID, err := strconv.ParseUint(parametros["retentorComandoID"], 10, 64)
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

	repositorio := repositories.NewRetentoresComandoRepository(db)
	if err = repositorio.Deletar(retentorComandoID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarRetentoresComandoPorNome busca todos os retentores de comando salvos no banco de dados pelo seu nome
func BuscarRetentoresComandoPorNome(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nome := parametros["nomeRetentorComando"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresComandoRepository(db)
	retentoresComando, err := repositorio.BuscarRetentoresComandoPorNome(nome)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, retentoresComando)
}
