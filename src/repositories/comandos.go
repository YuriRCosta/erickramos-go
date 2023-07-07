package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// Comandos representa um repositório de comandos
type Comandos struct {
	db *sql.DB
}

// NovoRepositorioDeComandos cria um repositório de comandos
func NovoRepositorioDeComandos(db *sql.DB) *Comandos {
	return &Comandos{db}
}

// Criar insere um comando no banco de dados
func (repositorio Comandos) Criar(comando models.Comando) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into comandos (nome, preco, qtd_estoque) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(comando.Nome, comando.Preco, comando.Qtd_estoque)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarComandos traz todos os comandos registrados no banco de dados
func (repositorio Comandos) BuscarComandos() ([]models.Comando, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from comandos",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var comandos []models.Comando

	for linhas.Next() {
		var comando models.Comando

		if err = linhas.Scan(
			&comando.ID,
			&comando.Nome,
			&comando.Preco,
			&comando.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		comandos = append(comandos, comando)
	}

	return comandos, nil
}

// BuscarComandoPorID traz um comando do banco de dados
func (repositorio Comandos) BuscarComandoPorID(ID uint64) (models.Comando, error) {
	linha, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from comandos where id = ?",
		ID,
	)
	if err != nil {
		return models.Comando{}, err
	}
	defer linha.Close()

	var comando models.Comando

	if linha.Next() {

		if err = linha.Scan(
			&comando.ID,
			&comando.Nome,
			&comando.Preco,
			&comando.Qtd_estoque,
		); err != nil {
			return models.Comando{}, err
		}
	}

	return comando, nil
}

// Atualizar altera as informações de um comando no banco de dados
func (repositorio Comandos) Atualizar(ID uint64, comando models.Comando) error {
	statement, err := repositorio.db.Prepare(
		"update comandos set nome = ?, preco = ?, qtd_estoque = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(comando.Nome, comando.Preco, comando.Qtd_estoque, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui um comando do banco de dados
func (repositorio Comandos) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from comandos where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarComandoPorNome traz comandos que contenham determinado nome
func (repositorio Comandos) BuscarComandoPorNome(nome string) ([]models.Comando, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque from comandos where nome like ?",
		"%"+nome+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var comandos []models.Comando

	for linhas.Next() {

		var comando models.Comando

		if err = linhas.Scan(
			&comando.ID,
			&comando.Nome,
			&comando.Preco,
			&comando.Qtd_estoque,
		); err != nil {
			return nil, err
		}

		comandos = append(comandos, comando)
	}

	return comandos, nil
}
