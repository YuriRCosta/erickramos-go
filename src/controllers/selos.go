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

// CriarSelo cria um selo no banco de dados
func CriarSelo(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var selo models.Selo
	if err = json.Unmarshal(corpoRequisicao, &selo); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	selo.ID, err = repositorio.Criar(selo)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, selo)
}

// BuscarTodosSelos traz todos os selos salvos no banco de dados
func BuscarTodosSelos(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	selos, err := repositorio.BuscarTodos()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, selos)
}

// BuscarSelos busca todos os selos salvos no banco de dados
func BuscarSelos(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	pagina, err := strconv.ParseInt(parametros["pagina"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

    // Definir um valor padrão para a página e itensPorPagina se não forem fornecidos ou inválidos
    if pagina <= 0 {
        pagina = 1
    }

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	selos, err := repositorio.BuscarSelos(pagina)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, selos)
}

// BuscarSeloPorID busca um selo salvo no banco de dados pelo seu ID
func BuscarSeloPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	seloID, err := strconv.ParseUint(parametros["seloID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeSelos(db)
	selo, err := repositorio.BuscarSeloPorID(seloID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, selo)
}

// AtualizarSelo atualiza um selo no banco de dados
func AtualizarSelo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	seloID, err := strconv.ParseUint(parametros["seloID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var selo models.Selo
	if err = json.Unmarshal(corpoRequisicao, &selo); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	if err = repositorio.Atualizar(seloID, selo); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarSelo deleta um selo do banco de dados
func DeletarSelo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	seloID, err := strconv.ParseUint(parametros["seloID"], 10, 64)
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

	repositorio := repositories.NovoRepositorioDeSelos(db)
	if err = repositorio.Deletar(seloID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeloPorNome busca selos salvos no banco de dados que contenham o nome informado
func BuscarSeloPorNome(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nomeSelo := parametros["nomeSelo"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	selo, err := repositorio.BuscarSeloPorNome(nomeSelo)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, selo)
}

// BuscarSeloPorMedida busca selos salvos no banco de dados que contenham a medida informada
func BuscarSeloPorMedida(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	medidaSelo := parametros["medidaSelo"]

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	selo, err := repositorio.BuscarSeloPorMedida(medidaSelo)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, selo)
}

// AdicionarEstoque adiciona uma quantidade de selos ao estoque
func AdicionarEstoque(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	seloID, err := strconv.ParseUint(parametros["seloID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var selo models.Selo
	if err = json.Unmarshal(corpoRequisicao, &selo); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeSelos(db)
	if err = repositorio.AdicionarEstoque(seloID, selo); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}