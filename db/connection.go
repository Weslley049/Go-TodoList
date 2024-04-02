package db

import (
	"database/sql"
	"fmt"

	confings "github.com/Weslley049/todo-list/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := confings.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		//Não utilizar em produção
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}
