package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// RetentoresComandoRepository representa um repositório de retentoresComando
type RetentoresComandoRepository struct {
	db *sql.DB
}

// NewRetentoresComandoRepository cria um repositório de retentoresComando
func NewRetentoresComandoRepository(db *sql.DB) *RetentoresComandoRepository {
	return &RetentoresComandoRepository{db}
}

// Criar insere um retentoresComando no banco de dados
func (repositorio RetentoresComandoRepository) Criar(retentoresComando models.RetentorComando) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into retentores_comandos (nome, preco, qtd_estoque) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(retentoresComando.Nome, retentoresComando.Preco, retentoresComando.Qtd_estoque)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarRetentoresComando traz todos os retentoresComando registrados no banco de dados
func (repositorio RetentoresComandoRepository) BuscarRetentoresComando() ([]models.RetentorComando, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from retentores_comandos",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var retentoresComando []models.RetentorComando

	for linhas.Next() {

		var retentorComando models.RetentorComando

		if err = linhas.Scan(
			&retentorComando.ID,
			&retentorComando.Nome,
			&retentorComando.Preco,
			&retentorComando.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		retentoresComando = append(retentoresComando, retentorComando)
	}

	return retentoresComando, nil
}

// BuscarRetentoresComandoPorID traz um retentoresComando do banco de dados
func (repositorio RetentoresComandoRepository) BuscarRetentorComandoPorID(ID uint64) (models.RetentorComando, error) {
	linha, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from retentores_comandos where id = ?",
		ID,
	)
	if err != nil {
		return models.RetentorComando{}, err
	}
	defer linha.Close()

	var retentorComando models.RetentorComando

	if linha.Next() {

		if err = linha.Scan(
			&retentorComando.ID,
			&retentorComando.Nome,
			&retentorComando.Preco,
			&retentorComando.Qtd_estoque,
		); err != nil {
			return models.RetentorComando{}, err
		}
	}

	return retentorComando, nil
}

// Atualizar altera as informações de um retentoresComando no banco de dados
func (repositorio RetentoresComandoRepository) Atualizar(ID uint64, retentoresComando models.RetentorComando) error {
	statement, err := repositorio.db.Prepare(
		"update retentores_comandos set nome = ?, preco = ?, qtd_estoque = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(retentoresComando.Nome, retentoresComando.Preco, retentoresComando.Qtd_estoque, ID)
	if err != nil {
		return err
	}

	return nil
}

// Deletar retira um retentoresComando do banco de dados
func (repositorio RetentoresComandoRepository) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from retentores_comandos where id = ?",
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

// BuscarRetentoresComandoPorNome traz todos os retentoresComando que atendem um nome específico
func (repositorio RetentoresComandoRepository) BuscarRetentoresComandoPorNome(nome string) ([]models.RetentorComando, error) {
	nome = "%" + nome + "%"

	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from retentores_comandos where nome like ?",
		nome,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var retentoresComando []models.RetentorComando

	for linhas.Next() {
		var retentorComando models.RetentorComando

		if err = linhas.Scan(
			&retentorComando.ID,
			&retentorComando.Nome,
			&retentorComando.Preco,
			&retentorComando.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		retentoresComando = append(retentoresComando, retentorComando)
	}

	return retentoresComando, nil
}
