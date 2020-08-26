package bd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func NuevaConexionBD() {
	//Se configura la conexion a la base de datos
	db, err = sql.Open("mysql", string("root:Int0603239I7!@tcp(127.0.0.1:3306)/hispanic"))
	revisarError(err)
	//Se comprueba que la conexion siga activa
	err = db.Ping()
	revisarError(err)
}

func TerminarConcexionBD() {
	//Se termina la conexion
	db.Close()
}

func revisarError(err error) {
	//Si hay algun problema cuando estemos trabajando con la BD nos lo hará saber en la consola
	if err != nil {
		panic(err)
	}
}
