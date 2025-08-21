package capadatos

import (
	"github.com/angorita/loft/db"
	"github.com/angorita/loft/models"
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
