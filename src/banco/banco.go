package banco

import (
	"database/sql"
	"erickramos-go/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
