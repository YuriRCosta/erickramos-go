package repositories

import (
	"database/sql"
	"erickramos-go/src/models"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = "%" + nomeOuNick + "%"

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna seu ID e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

// Seguir permite que um usuário siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

// PararDeSeguir permite que um usuário pare de seguir outro
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var seguidores []models.Usuario

	for linhas.Next() {
		var seguidor models.Usuario

		if err = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); err != nil {
			return nil, err
		}

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?
	`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var seguindo []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		seguindo = append(seguindo, usuario)
	}

	return seguindo, nil
}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	line, err := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var usuario models.Usuario

	if line.Next() {
		if err = line.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senhaComHash string) error {
	statement, err := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(senhaComHash, usuarioID); err != nil {
		return err
	}

	return nil
}
