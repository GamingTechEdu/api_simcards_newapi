package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDB *sql.DB

func ConnectMysql() {
	var err error
	db, err := sql.Open("mysql", "kdl:sim$srv%8484@tcp(187.122.102.36:50026)/KDL_MANUTENCAO")
	db.SetMaxIdleConns(100)

	if err != nil {
		fmt.Println("Error Init", err.Error())
	}

	fmt.Println("Connected to Mysql")
	MysqlDB = db
}
