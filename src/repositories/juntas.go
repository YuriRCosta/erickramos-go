package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// Juntas representa um repositório de juntas
type Juntas struct {
	db *sql.DB
}

// NovoRepositorioDeJuntas cria um repositório de juntas
func NovoRepositorioDeJuntas(db *sql.DB) *Juntas {
	return &Juntas{db}
}

// Criar insere um junta no banco de dados
func (repositorio Juntas) Criar(junta models.Junta) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into juntas (cabecotes, preco, qtd_estoque) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(junta.Cabecotes, junta.Preco, junta.Qtd_estoque)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarJuntas traz todos os juntas registrados no banco de dados
func (repositorio Juntas) BuscarJuntas() ([]models.Junta, error) {
	linhas, err := repositorio.db.Query(
		"select id, cabecotes, preco, qtd_estoque from juntas",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var juntas []models.Junta

	for linhas.Next() {
		var junta models.Junta

		if err = linhas.Scan(
			&junta.ID,
			&junta.Cabecotes,
			&junta.Preco,
			&junta.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		juntas = append(juntas, junta)
	}

	return juntas, nil
}

// BuscarJuntaPorID traz um junta do banco de dados
func (repositorio Juntas) BuscarJuntaPorID(ID uint64) (models.Junta, error) {
	linha, err := repositorio.db.Query(
		"select id, cabecotes, preco, qtd_estoque from juntas where id = ?",
		ID,
	)
	if err != nil {
		return models.Junta{}, err
	}
	defer linha.Close()

	var junta models.Junta

	if linha.Next() {

		if err = linha.Scan(
			&junta.ID,
			&junta.Cabecotes,
			&junta.Preco,
			&junta.Qtd_estoque,
		); err != nil {
			return models.Junta{}, err
		}
	}

	return junta, nil
}

// Atualizar altera as informações de um junta no banco de dados
func (repositorio Juntas) Atualizar(ID uint64, junta models.Junta) error {
	statement, err := repositorio.db.Prepare(
		"update juntas set cabecotes = ?, preco = ?, qtd_estoque = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(junta.Cabecotes, junta.Preco, junta.Qtd_estoque, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui um junta do banco de dados
func (repositorio Juntas) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from juntas where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarJuntaPorCabecotes traz juntas que contenham determinado cabecotes
func (repositorio Juntas) BuscarJuntaPorCabecotes(cabecotes string) ([]models.Junta, error) {
	linhas, err := repositorio.db.Query(
		"select id, cabecotes, preco, qtd_estoque from juntas where cabecotes like ?",
		"%"+cabecotes+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var juntas []models.Junta

	for linhas.Next() {

		var junta models.Junta

		if err = linhas.Scan(
			&junta.ID,
			&junta.Cabecotes,
			&junta.Preco,
			&junta.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		juntas = append(juntas, junta)
	}

	return juntas, nil
}
