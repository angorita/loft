package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var cn *sql.DB

func Open() {
	/*
		Dia 18 de octubre 2024 me salio la coneccion con sqlite yeahhhhhh!!
	*/
	db, err := sql.Open("sqlite3", "db/loft.db")
	if err != nil {
		panic(err)
	}
	cn = db
}
func Query(consulta string, args ...interface{}) (*sql.Rows, error) {
	rows, err := cn.Query(consulta, args...)
	if err != nil {
		panic("No se pudo hacer la consulta....")
	}
	return rows, err
}
func Exec(consulta string, args ...interface{}) (sql.Result, error) {
	result, err := cn.Exec(consulta, args...)
	if err != nil {
		return nil, errors.New("no se pudo realizar la consulta")
	}
	return result, err
}
func Close() {
	cn.Close()
}
