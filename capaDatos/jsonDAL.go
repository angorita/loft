package capadatos

import (
	"LoftSQLite3/db"
	"LoftSQLite3/models"
)

func Wilder() models.ListaWilder {
	oLista := models.ListaWilder{}
	sqlQuery := `select descripcion,precio,cantidad
	from producto where descripcion like '%ilder%'`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oJson := models.Wilder{}
		rows.Scan(&oJson.Descripcion, &oJson.Precio, &oJson.Cantidad)
		oLista = append(oLista, oJson)
	}
	return oLista
}
