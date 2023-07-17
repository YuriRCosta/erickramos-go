package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// Valvulas representa um repositório de valvulas
type Valvulas struct {
	db *sql.DB
}

// NovoRepositorioDeValvulas cria um repositório de valvulas
func NovoRepositorioDeValvulas(db *sql.DB) *Valvulas {
	return &Valvulas{db}
}

// Criar insere um valvula no banco de dados
func (repositorio Valvulas) Criar(valvula models.Valvula) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into valvulas (nome, codigo, preco, qtd_estoque, tipo) values (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(valvula.Nome, valvula.Codigo, valvula.Preco, valvula.Qtd_estoque, valvula.Tipo)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarValvulas traz todos os valvulas registrados no banco de dados
func (repositorio Valvulas) BuscarValvulas() ([]models.Valvula, error) {
	linhas, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque, tipo from valvulas",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var valvulas []models.Valvula

	for linhas.Next() {
		var valvula models.Valvula

		if err = linhas.Scan(
			&valvula.ID,
			&valvula.Codigo,
			&valvula.Nome,
			&valvula.Preco,
			&valvula.Qtd_estoque,
			&valvula.Tipo,
		); err != nil {
			return nil, err
		}

		valvulas = append(valvulas, valvula)
	}

	return valvulas, nil
}

// BuscarValvulasPorPagina traz todos os valvulas registrados no banco de dados de acordo com a paginação
func (repositorio Valvulas) BuscarValvulasPorPagina(pagina uint64) ([]models.Valvula, error) {
	offset := (pagina - 1) * 5
	linhas, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque, tipo from valvulas limit 5 offset ?",
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var valvulas []models.Valvula

	for linhas.Next() {
		var valvula models.Valvula

		if err = linhas.Scan(
			&valvula.ID,
			&valvula.Codigo,
			&valvula.Nome,
			&valvula.Preco,
			&valvula.Qtd_estoque,
			&valvula.Tipo,
		); err != nil {
			return nil, err
		}

		valvulas = append(valvulas, valvula)
	}

	return valvulas, nil
}

// BuscarValvulaPorID traz um valvula do banco de dados
func (repositorio Valvulas) BuscarValvulaPorID(ID uint64) (models.Valvula, error) {
	linha, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque, tipo from valvulas where id = ?",
		ID,
	)
	if err != nil {
		return models.Valvula{}, err
	}
	defer linha.Close()

	var valvula models.Valvula

	if linha.Next() {

		if err = linha.Scan(
			&valvula.ID,
			&valvula.Codigo,
			&valvula.Nome,
			&valvula.Preco,
			&valvula.Qtd_estoque,
			&valvula.Tipo,
		); err != nil {
			return models.Valvula{}, err
		}
	}

	return valvula, nil
}

// Atualizar altera as informações de um valvula no banco de dados
func (repositorio Valvulas) Atualizar(ID uint64, valvula models.Valvula) error {
	statement, err := repositorio.db.Prepare(
		"update valvulas set codigo = ?, nome = ?, preco = ?, qtd_estoque = ?, tipo = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(valvula.Codigo, valvula.Nome, valvula.Preco, valvula.Qtd_estoque, valvula.Tipo, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui um valvula do banco de dados
func (repositorio Valvulas) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from valvulas where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarValvulaPorNome traz valvulas que contenham determinado nome
func (repositorio Valvulas) BuscarValvulaPorNome(nome string) ([]models.Valvula, error) {
	linhas, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque, tipo from valvulas where nome like ?",
		"%"+nome+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var valvulas []models.Valvula

	for linhas.Next() {

		var valvula models.Valvula

		if err = linhas.Scan(
			&valvula.ID,
			&valvula.Codigo,
			&valvula.Nome,
			&valvula.Preco,
			&valvula.Qtd_estoque,
			&valvula.Tipo,
		); err != nil {
			return nil, err
		}

		valvulas = append(valvulas, valvula)
	}

	return valvulas, nil
}

// BuscarValvulaPorTipo traz valvulas que contenham determinada tipo
func (repositorio Valvulas) BuscarValvulaPorTipo(tipo string) ([]models.Valvula, error) {
	linhas, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque, tipo from valvulas where tipo like ?",
		"%"+tipo+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var valvulas []models.Valvula

	for linhas.Next() {

		var valvula models.Valvula

		if err = linhas.Scan(
			&valvula.ID,
			&valvula.Codigo,
			&valvula.Nome,
			&valvula.Preco,
			&valvula.Qtd_estoque,
			&valvula.Tipo,
		); err != nil {
			return nil, err
		}

		valvulas = append(valvulas, valvula)
	}

	return valvulas, nil
}

// AdicionarEstoqueValvula adiciona uma determinada quantidade de estoque a um valvula
func (repositorio Valvulas) AdicionarEstoqueValvula(ID uint64, quantidade uint64) error {
	statement, err := repositorio.db.Prepare("update valvulas set qtd_estoque = qtd_estoque + ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(quantidade, ID); err != nil {
		return err
	}

	return nil
}