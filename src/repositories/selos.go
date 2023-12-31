package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// Selos representa um repositório de selos
type Selos struct {
	db *sql.DB
}

// NovoRepositorioDeSelos cria um repositório de selos
func NovoRepositorioDeSelos(db *sql.DB) *Selos {
	return &Selos{db}
}

// Criar insere um selo no banco de dados
func (repositorio Selos) Criar(selo models.Selo) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into selos (nome, preco, qtd_estoque, medida) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(selo.Nome, selo.Preco, selo.Qtd_estoque, selo.Medida)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarTodos traz todos os selos registrados no banco de dados
func (repositorio Selos) BuscarTodos() ([]models.Selo, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque, medida from selos order by nome",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var selos []models.Selo

	for linhas.Next() {
		var selo models.Selo

		if err = linhas.Scan(
			&selo.ID,
			&selo.Nome,
			&selo.Preco,
			&selo.Qtd_estoque,
			&selo.Medida,
		); err != nil {
			return nil, err
		}

		selos = append(selos, selo)
	}

	return selos, nil
}

// BuscarSelos traz todos os selos registrados no banco de dados
func (repositorio Selos) BuscarSelos(pagina int64) ([]models.Selo, error) {
	offset := (pagina - 1) * 5

	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque, medida from selos order by nome LIMIT 5 OFFSET ?",
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var selos []models.Selo

	for linhas.Next() {
		var selo models.Selo

		if err = linhas.Scan(
			&selo.ID,
			&selo.Nome,
			&selo.Preco,
			&selo.Qtd_estoque,
			&selo.Medida,
		); err != nil {
			return nil, err
		}

		selos = append(selos, selo)
	}

	return selos, nil
}

// BuscarSeloPorID traz um selo do banco de dados
func (repositorio Selos) BuscarSeloPorID(ID uint64) (models.Selo, error) {
	linha, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque, medida from selos where id = ?",
		ID,
	)
	if err != nil {
		return models.Selo{}, err
	}
	defer linha.Close()

	var selo models.Selo

	if linha.Next() {

		if err = linha.Scan(
			&selo.ID,
			&selo.Nome,
			&selo.Preco,
			&selo.Qtd_estoque,
			&selo.Medida,
		); err != nil {
			return models.Selo{}, err
		}
	}

	return selo, nil
}

// Atualizar altera as informações de um selo no banco de dados
func (repositorio Selos) Atualizar(ID uint64, selo models.Selo) error {
	statement, err := repositorio.db.Prepare(
		"update selos set nome = ?, preco = ?, qtd_estoque = ?, medida = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(selo.Nome, selo.Preco, selo.Qtd_estoque, selo.Medida, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui um selo do banco de dados
func (repositorio Selos) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from selos where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarSeloPorNome traz selos que contenham determinado nome
func (repositorio Selos) BuscarSeloPorNome(nome string) ([]models.Selo, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque, medida from selos where nome like ?",
		"%"+nome+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var selos []models.Selo

	for linhas.Next() {

		var selo models.Selo

		if err = linhas.Scan(
			&selo.ID,
			&selo.Nome,
			&selo.Preco,
			&selo.Qtd_estoque,
			&selo.Medida,
		); err != nil {
			return nil, err
		}

		selos = append(selos, selo)
	}

	return selos, nil
}

// BuscarSeloPorMedida traz selos que contenham determinada medida
func (repositorio Selos) BuscarSeloPorMedida(medida string) ([]models.Selo, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, preco, qtd_estoque, medida from selos where medida like ?",
		"%"+medida+"%",
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var selos []models.Selo

	for linhas.Next() {

		var selo models.Selo

		if err = linhas.Scan(
			&selo.ID,
			&selo.Nome,
			&selo.Preco,
			&selo.Qtd_estoque,
			&selo.Medida,
		); err != nil {
			return nil, err
		}

		selos = append(selos, selo)
	}

	return selos, nil
}

// BuscarSelosPaginacao busca selos no banco de dados com paginação
func (repositorio Selos) BuscarSelosPaginacao(pagina int) ([]models.Selo, error) {
	linhas, err := repositorio.db.Query(
		"select * from selos limit 5 offset ?",
		pagina,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var selos []models.Selo

	for linhas.Next() {

		var selo models.Selo

		if err = linhas.Scan(
			&selo.ID,
			&selo.Nome,
			&selo.Preco,
			&selo.Qtd_estoque,
			&selo.Medida,
		); err != nil {
			return nil, err
		}

		selos = append(selos, selo)
	}

	return selos, nil
}


// AdicionarEstoque adiciona a quantidade de selos em estoque
func (repositorio Selos) AdicionarEstoque(ID uint64, selo models.Selo) error {
	statement, err := repositorio.db.Prepare(
		"update selos set qtd_estoque = qtd_estoque + ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(selo.Qtd_estoque, ID); err != nil {
		return err
	}

	return nil
}