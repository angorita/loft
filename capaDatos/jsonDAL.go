package capadatos

import (
	"X/db"
	m "X/models"
)

func Wilder() m.ListaWilder {
	oLista := m.ListaWilder{}
	sqlQuery := `select descripcion,precio,cantidad
	from producto where descripcion like '%ilder%'`

	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oJson := m.Wilder{}
		rows.Scan(&oJson.Descripcion, &oJson.Precio, &oJson.Cantidad)
		oLista = append(oLista, oJson)
	}
	return oLista
}
