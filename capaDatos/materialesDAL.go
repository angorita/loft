// https://www.guru99.com/es/sqlite-query.html
package capadatos

import (
	db "LoftSQLite3/db"
	m "LoftSQLite3/models"
	"database/sql"
)

/*          $total = $producto['precio'] * $producto['cantidad'];
            $parcial += $total;
            $totalUsa = ($producto['precio'] * $producto['cantidad'] ) / $producto['dolar'];
            $parcialUsa += $totalUsa;

*/
// LISTA SIMPLE
func ListaMateriales() m.ListaMateriales {
	oListaMateriales := m.ListaMateriales{}
	sqlQuery := `SELECT id,
	descripcion,precio,cantidad,
	precio*cantidad as total,
	strftime('%m-%Y',fecha),
	dolar,
	bhabilitado=1 
	FROM producto
	Order by id asc
	` //consulta que funciona con sqlite3
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

// LISTAS CON PARAMETROS NOMBREMATERIALES
// chequear numero de campos en base de datos y en funciones...
func FiltrarMateriales(nombreMat string) m.ListaMateriales {
	oListaMateriales := m.ListaMateriales{}
	sqlQuery := `SELECT * FROM  producto
	where descripcion like '%'||$1||'%'`
	db.Open()
	rows, _ := db.Query(sqlQuery, nombreMat)
	for rows.Next() {
		oMateriales := m.Materiales{}
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio,
			&oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar,
			&oMateriales.Bhabilitado)
		oListaMateriales = append(oListaMateriales, oMateriales)

	}
	return oListaMateriales
}

// LISTAS COMBOBOX CON PARAMETRO ID
func FiltrarId(IdMaterial int) m.ListaMateriales {
	oComboMateriales := m.ListaMateriales{}
	sqlQuery := `select * from producto where id = ( $1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, IdMaterial)
	for rows.Next() {
		oMateriales := m.Materiales{}
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio,
			&oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar,
			&oMateriales.Bhabilitado)
		oComboMateriales = append(oComboMateriales, oMateriales)
	}
	db.Close()
	return oComboMateriales
}
func BuscarMaterialesPorId(id int) m.Materiales {

	sqlQuery := `select descripcion,precio,cantidad,fecha,dolar,id from producto where id=($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, id)
	oMateriales := m.Materiales{}

	for rows.Next() {
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio, &oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar)
	}
	db.Close()

	return oMateriales
}
func InsertarMaterial(descripcion string, precio float64, cantidad int, fecha string, dolar float64, bhabilitado bool) (sql.Result, error) {
	db.Open()
	sql := `insert into producto(descripcion, precio, cantidad, fecha, dolar, bhabilitado)values($1,$2,$3,$4,$5,$6)`
	//validacion
	// var errorMaterial error
	errorMaterial := m.Max(descripcion)
	if errorMaterial != nil {
		return nil, errorMaterial
	}
	result, err := db.Exec(sql, descripcion, precio, cantidad, fecha, dolar, true)
	db.Close()
	return result, err
}

// Cuando falta una columna se cuelga sin decir nada
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
