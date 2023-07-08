package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// RetentoresValvulaRepository representa um repositório de retentoresValvula
type RetentoresValvulaRepository struct {
	db *sql.DB
}

// NewRetentoresValvulaRepository cria um repositório de retentoresValvula
func NewRetentoresValvulaRepository(db *sql.DB) *RetentoresValvulaRepository {
	return &RetentoresValvulaRepository{db}
}

// Criar insere um retentoresValvula no banco de dados
func (repositorio RetentoresValvulaRepository) Criar(retentoresValvula models.RetentorValvula) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into retentores_valvulas (nome, preco, qtd_estoque) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(retentoresValvula.Nome, retentoresValvula.Preco, retentoresValvula.Qtd_estoque)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarRetentoresValvula traz todos os retentoresValvula registrados no banco de dados
func (repositorio RetentoresValvulaRepository) BuscarRetentoresValvula() ([]models.RetentorValvula, error) {
	linhas, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque from retentores_valvulas",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var retentoresValvula []models.RetentorValvula

	for linhas.Next() {

		var retentorValvula models.RetentorValvula

		if err = linhas.Scan(
			&retentorValvula.ID,
			&retentorValvula.Codigo,
			&retentorValvula.Nome,
			&retentorValvula.Preco,
			&retentorValvula.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		retentoresValvula = append(retentoresValvula, retentorValvula)
	}

	return retentoresValvula, nil
}

// BuscarRetentoresValvulaPorID traz um retentoresValvula do banco de dados
func (repositorio RetentoresValvulaRepository) BuscarRetentorValvulaPorID(ID uint64) (models.RetentorValvula, error) {
	linha, err := repositorio.db.Query(
		"select id, codigo, nome, preco, qtd_estoque from retentores_valvulas where id = ?",
		ID,
	)
	if err != nil {
		return models.RetentorValvula{}, err
	}
	defer linha.Close()

	var retentorValvula models.RetentorValvula

	if linha.Next() {

		if err = linha.Scan(
			&retentorValvula.ID,
			&retentorValvula.Codigo,
			&retentorValvula.Nome,
			&retentorValvula.Preco,
			&retentorValvula.Qtd_estoque,
		); err != nil {
			return models.RetentorValvula{}, err
		}
	}

	return retentorValvula, nil
}

// Atualizar altera as informações de um retentoresValvula no banco de dados
func (repositorio RetentoresValvulaRepository) Atualizar(ID uint64, retentoresValvula models.RetentorValvula) error {
	statement, err := repositorio.db.Prepare(
		"update retentores_valvulas set nome = ?, preco = ?, qtd_estoque = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(retentoresValvula.Nome, retentoresValvula.Preco, retentoresValvula.Qtd_estoque, ID)
	if err != nil {
		return err
	}

	return nil
}

// Deletar retira um retentoresValvula do banco de dados
func (repositorio RetentoresValvulaRepository) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from retentores_valvulas where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarRetentoresValvulaPorNome traz todos os retentoresValvula que atendem um nome específico
func (repositorio RetentoresValvulaRepository) BuscarRetentoresValvulaPorNome(nome string) ([]models.RetentorValvula, error) {
	nome = "%" + nome + "%"

	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from retentores_valvulas where nome like ?",
		nome,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var retentoresValvula []models.RetentorValvula

	for linhas.Next() {
		var retentorValvula models.RetentorValvula

		if err = linhas.Scan(
			&retentorValvula.ID,
			&retentorValvula.Nome,
			&retentorValvula.Preco,
			&retentorValvula.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		retentoresValvula = append(retentoresValvula, retentorValvula)
	}

	return retentoresValvula, nil
}
