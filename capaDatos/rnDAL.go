package capadatos

import (
	"X/db"
	m "X/models"
)

func ListaRN() m.ListaRN {
	oLista := m.ListaRN{}
	sqlQuery := `SELECT id,nom,am,ap,fn,parto,obs,neo,os from partos limit 1`
	db.Open()
	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic("No se puedo consultar ")
	}
	for rows.Next() {
		oRN := m.RN{}
		rows.Scan(&oRN.ID, &oRN.Nom, &oRN.AP, &oRN.AM, &oRN.FN, &oRN.Parto, &oRN.Obs, &oRN.Neo, &oRN.OS)
		oLista = append(oLista, oRN)
	}
	defer db.Close()
	return oLista
}
