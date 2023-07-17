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

// CriarRetentorValvula cria um retentor de valvula no banco de dados
func CriarRetentorValvula(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var retentorValvula models.RetentorValvula
	if err = json.Unmarshal(corpoRequisicao, &retentorValvula); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	retentorValvula.ID, err = repositorio.Criar(retentorValvula)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, retentorValvula)
}

// BuscarRetentoresValvula busca todos os retentores de valvula salvos no banco de dados
func BuscarRetentoresValvula(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	retentoresValvula, err := repositorio.BuscarRetentoresValvula()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, retentoresValvula)
}

// BuscarRetentoresValvulaPaginacao busca todos os retentores de valvula salvos no banco de dados com paginação
func BuscarRetentoresValvulaPaginacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	pagina, err := strconv.ParseInt(parametros["pagina"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

    if pagina <= 0 {
        pagina = 1
    }

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	retentoresValvula, err := repositorio.BuscarRetentoresValvulaPaginacao(pagina)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, retentoresValvula)
}

// BuscarRetentorValvulaPorID busca um retentor de valvula salvo no banco de dados pelo seu ID
func BuscarRetentorValvulaPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorValvulaID, err := strconv.ParseUint(parametros["retentorValvulaID"], 10, 64)
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

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	retentorValvula, err := repositorio.BuscarRetentorValvulaPorID(retentorValvulaID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, retentorValvula)
}

// AtualizarRetentorValvula atualiza um retentor de valvula no banco de dados
func AtualizarRetentorValvula(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorValvulaID, err := strconv.ParseUint(parametros["retentorValvulaID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var retentorValvula models.RetentorValvula
	if err = json.Unmarshal(corpoRequisicao, &retentorValvula); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	if err = repositorio.Atualizar(retentorValvulaID, retentorValvula); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarRetentorValvula deleta um retentor de valvula do banco de dados
func DeletarRetentorValvula(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorValvulaID, err := strconv.ParseUint(parametros["retentorValvulaID"], 10, 64)
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

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	if err = repositorio.Deletar(retentorValvulaID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarRetentoresValvulaPorNome busca todos os retentores de valvula salvos no banco de dados pelo seu nome
func BuscarRetentoresValvulaPorNome(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nome := parametros["nomeRetentorValvula"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NewRetentoresValvulaRepository(db)
	retentoresValvula, err := repositorio.BuscarRetentoresValvulaPorNome(nome)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, retentoresValvula)
}

// AdicionarEstoqueRetentorValvula adiciona uma quantidade X de itens ao estoque do retentor de valvula
func AdicionarEstoqueRetentorValvula(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	retentorValvulaID, err := strconv.ParseUint(parametros["retentorValvulaID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	var quantidadeEstoque models.RetentorValvula
	if err = json.Unmarshal(corpoRequisicao, &quantidadeEstoque); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	retentorValvulaRepository := repositories.NewRetentoresValvulaRepository(db)
	if err = retentorValvulaRepository.AdicionarEstoque(retentorValvulaID, uint64(quantidadeEstoque.Qtd_estoque)); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

