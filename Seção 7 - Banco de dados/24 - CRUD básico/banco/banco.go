package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // DRIVER DE CONEXÃO COM O MYSQL
)

func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
