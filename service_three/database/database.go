package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mailru/dbr"
	"service_three/globals"
)

var DbConnection *dbr.Session

func Init() {
	conn, err := dbr.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
		globals.Conf.DatabaseUser,
		globals.Conf.DatabasePassword,
		globals.Conf.DatabasePort,
		globals.Conf.DatabaseName), nil)
	if err != nil {
		panic(err)
	}

	DbConnection = conn.NewSession(nil)

	err = runMigrations()
	if err != nil {
		panic(err)
	}

}
