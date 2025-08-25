package capadatos

import (
	"database/sql"

	db "github.com/angorita/loft/db"
	m "github.com/angorita/loft/models"
)

func ListaMateriales() m.ListaMateriales {
	oListaMateriales := m.ListaMateriales{}
	sqlQuery := `SELECT id,
	descripcion,precio,cantidad,
	precio*cantidad as total,
	strftime('%m-%Y',fecha),
	dolar,
	bhabilitado
	FROM producto
	Order by id asc
	`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oMateriales := m.Materiales{}

		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio,
			&oMateriales.Cantidad, &oMateriales.Total,
			&oMateriales.Fecha, &oMateriales.Dolar, &oMateriales.Bhabilitado)
		oListaMateriales = append(oListaMateriales, oMateriales)
	}
	return oListaMateriales
}

// filtrado de materiales por descripcion...
func FiltrarMateriales(nombreMat string) m.ListaMateriales {
	oListaMateriales := m.ListaMateriales{}
	sqlQuery := `SELECT id,descripcion,precio,cantidad,strftime('%m-%Y',fecha)fecha,dolar FROM  producto
	where descripcion like '%'||$1||'%'`
	db.Open()
	rows, _ := db.Query(sqlQuery, nombreMat)
	for rows.Next() {
		oMateriales := m.Materiales{}
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio,
			&oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar)
		oListaMateriales = append(oListaMateriales, oMateriales)

	}
	return oListaMateriales
}

// error en el orden de los campos en sqlquery y por ende en rows.Scan
func FiltrarId(IdMaterial int) m.ListaMateriales {
	oComboMateriales := m.ListaMateriales{}
	sqlQuery := `select id,descripcion,precio,cantidad,
	fecha,dolar from producto where id = ( $1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, IdMaterial)
	for rows.Next() {
		oMateriales := m.Materiales{}
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio,
			&oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar,
		)
		oComboMateriales = append(oComboMateriales, oMateriales)
	}
	db.Close()
	return oComboMateriales
}
func BuscarMaterialesPorId(id int) m.Materiales {
	//el orden del query debe ser respetado en rows...
	sqlQuery := `select id,descripcion,precio,cantidad,fecha,dolar from producto where id=($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, id)
	oMateriales := m.Materiales{}

	for rows.Next() {
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio, &oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar)
	}
	db.Close()

	return oMateriales
}
func InsertarMaterial(descripcion string, precio float64, cantidad int, fecha string, dolar float64, bhabilitado int) (sql.Result, error) {
	db.Open()
	sql := `insert into producto(descripcion, precio, cantidad, fecha, dolar, bhabilitado)values($1,$2,$3,$4,$5,$6)`
	errorMaterial := m.Max(descripcion)
	if errorMaterial != nil {
		return nil, errorMaterial
	}
	result, err := db.Exec(sql, descripcion, precio, cantidad, fecha, dolar, bhabilitado)
	db.Close()
	return result, err
}

func Actualizar(id int, descripcion string, precio float64, cantidad int, fecha string, dolar float64) (sql.Result, error) {
	db.Open()
	sql := `update producto set id=$1,descripcion=$2,precio=$3,cantidad=$4,fecha=$5,dolar=$6 where id=$1`
	result, err := db.Exec(sql, id, descripcion, precio, cantidad, fecha, dolar)
	db.Close()
	return result, err
}
func Eliminar(id int) (sql.Result, error) {
	db.Open()
	sql := `delete from producto where id=$1`
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
