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
		"insert into users (full_name, user_name, password, enabled, credentials_non_expired, account_non_locked, account_non_expired) values (?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.FullName, usuario.Username, usuario.Password, usuario.Enabled, usuario.CredentialsNonExpired, usuario.AccountNonLocked, usuario.AccountNonExpired)
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
func (repositorio Usuarios) Buscar(nome string) ([]models.Usuario, error) {
	nome = "%" + nome + "%"

	linhas, err := repositorio.db.Query(
		"select id, full_name, user_name from users where full_name LIKE ?",
		nome,
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
			&usuario.FullName,
			&usuario.Username,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
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

// BuscarPorUsername busca um usuário por username e retorna seu ID e senha com hash
func (repositorio Usuarios) BuscarPorUsername(username string) (models.Usuario, error) {
	linha, err := repositorio.db.Query("select id, password from users where user_name = ?", username)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Password); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	line, err := repositorio.db.Query("select password from users where id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var usuario models.Usuario

	if line.Next() {
		if err = line.Scan(&usuario.Password); err != nil {
			return "", err
		}
	}

	return usuario.Password, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senhaComHash string) error {
	statement, err := repositorio.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(senhaComHash, usuarioID); err != nil {
		return err
	}

	return nil
}
