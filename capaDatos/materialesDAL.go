package capadatos

import (
	"X/db"
	m "X/models"
	"database/sql"
)

func ListaMateriales() m.ListaMateriales {
	oListaMateriales := m.ListaMateriales{}
	sqlQuery := `SELECT id,
	descripcion,precio,cantidad,
	precio*cantidad as total,
	strftime('%d-%m-%Y',fecha),
	dolar,
	bhabilitado
	FROM producto 
	where bhabilitado=1
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
	db.Close()
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
	db.Close()
	return oListaMateriales
}

// error en el orden de los campos en sqlquery y por ende en rows.Scan
func FiltrarId(IdMaterial int) m.ListaMateriales {
	oComboMateriales := m.ListaMateriales{}
	sqlQuery := `
	select id,descripcion,precio,cantidad,
	fecha,dolar 
	from producto 
	where id = ( $1)
	`
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
	sqlQuery := `
	select id,descripcion,precio,cantidad,fecha,dolar 
	from producto 
	where id=($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, id)
	oMateriales := m.Materiales{}

	for rows.Next() {
		rows.Scan(&oMateriales.Id, &oMateriales.Descripcion, &oMateriales.Precio,
			&oMateriales.Cantidad, &oMateriales.Fecha, &oMateriales.Dolar)
	}
	db.Close()
	return oMateriales
}

/*
VALUES (valor1, valor2, ...):
Aquí proporcionas los valores que corresponden a cada columna listada.
Los valores deben estar en el mismo orden que las columnas.
INSERT INTO productos (nombre, precio) VALUES ('Laptop', 1200.00);
id se incluye automaticamente.si es una clave primaria autoincremental.
*/
func InsertarMaterial(descripcion string, precio float64, cantidad int, fecha string, dolar float64, bhabilitado int) (sql.Result, error) {
	db.Open()
	sql := `
	insert 
	into producto(descripcion, precio, cantidad, fecha, dolar, bhabilitado)
	values($1,$2,$3,$4,$5,$6)`
	errorMaterial := m.Max(descripcion)
	if errorMaterial != nil {
		return nil, errorMaterial
	}
	result, err := db.Exec(sql, descripcion, precio, cantidad,
		fecha, dolar, bhabilitado)
	db.Close()
	return result, err
}

/*
El error estaba en el exec no tenia el mismo orden que la consulta sql
update producto set descripcion=$1,precio=$2,cantidad=$3,fecha=$4,dolar=$5 where id=$6
db.Exec(sql, descripcion, precio, cantidad, fecha, dolar, id)
el id estaba primero en exec, lo cambie, y se pudo guardar.
tarde mas o menos 2 meses ....la puta que te parioooooo!!!!
tuve que poner fmt's en todos lados y me salioooooo!!!!
hoy 4 de septiembre de 2025
					orden
					*********
1 func parametros 	(1,2,3,4)
2 sql 				(1,2,3,4)
3 exec 				(1,2,3,4)
*/

func Actualizar(descripcion string, precio float64, cantidad int,
	fecha string, dolar float64, id int) (sql.Result, error) {
	db.Open()
	sql := `
	update producto 
	set descripcion=$1,
	precio=$2,cantidad=$3,
	fecha=$4,dolar=$5 
	where id=$6`
	result, err := db.Exec(sql, descripcion, precio, cantidad, fecha, dolar, id)
	db.Close()
	return result, err
}

func Eliminar(id int) (sql.Result, error) {
	db.Open()
	sql := `update producto set bhabilitado=0 where id=$1`
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
