package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com mysql
)

//Abre a conxexao com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:mE1@golangaz@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
