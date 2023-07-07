package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// Cabecotes representa um repositório de cabecotes
type Cabecotes struct {
	db *sql.DB
}

// NovoRepositorioDeCabecotes cria um repositório de cabecotes
func NovoRepositorioDeCabecotes(db *sql.DB) *Cabecotes {
	return &Cabecotes{db}
}

// Criar insere um cabecote no banco de dados
func (repositorio Cabecotes) Criar(cabecote models.Cabecote) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into cabecotes (nome, material, qtd_valvulas) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(cabecote.Nome, cabecote.Material, cabecote.Qtd_valvulas)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarCabecotes traz todos os cabecotes registrados no banco de dados
func (repositorio Cabecotes) BuscarCabecotes() ([]models.Cabecote, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, qtd_valvulas, material from cabecotes",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var cabecotes []models.Cabecote

	for linhas.Next() {
		var cabecote models.Cabecote

		if err = linhas.Scan(
			&cabecote.ID,
			&cabecote.Nome,
			&cabecote.Qtd_valvulas,
			&cabecote.Material,
		); err != nil {
			return nil, err
		}

		cabecotes = append(cabecotes, cabecote)
	}

	return cabecotes, nil
}

// BuscarCabecotePorID traz um cabecote do banco de dados
func (repositorio Cabecotes) BuscarCabecotePorID(ID uint64) (models.Cabecote, error) {
	linha, err := repositorio.db.Query(
		"select id, nome, material, qtd_valvulas from cabecotes where id = ?",
		ID,
	)
	if err != nil {
		return models.Cabecote{}, err
	}
	defer linha.Close()

	var cabecote models.Cabecote

	if linha.Next() {

		if err = linha.Scan(
			&cabecote.ID,
			&cabecote.Nome,
			&cabecote.Material,
			&cabecote.Qtd_valvulas,
		); err != nil {
			return models.Cabecote{}, err
		}
	}

	return cabecote, nil
}

// Atualizar altera as informações de um cabecote no banco de dados
func (repositorio Cabecotes) Atualizar(ID uint64, cabecote models.Cabecote) error {
	statement, err := repositorio.db.Prepare(
		"update cabecotes set nome = ?, material = ?, qtd_valvulas = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(cabecote.Nome, cabecote.Material, cabecote.Qtd_valvulas, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui um cabecote do banco de dados
func (repositorio Cabecotes) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from cabecotes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarCabecotePorNome traz cabecotes que contenham determinado nome
func (repositorio Cabecotes) BuscarCabecotePorNome(nome string) ([]models.Cabecote, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, material, qtd_valvulas from cabecotes where nome like ?",
		"%"+nome+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var cabecotes []models.Cabecote

	for linhas.Next() {

		var cabecote models.Cabecote

		if err = linhas.Scan(
			&cabecote.ID,
			&cabecote.Nome,
			&cabecote.Material,
			&cabecote.Qtd_valvulas,
		); err != nil {
			return nil, err
		}

		cabecotes = append(cabecotes, cabecote)
	}

	return cabecotes, nil
}
