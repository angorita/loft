package handlers

import (
	cd "X/capaDatos"
	m "X/models"
	u "X/utilitarios"
	"net/http"
	"strconv"
)

/*
Quien llama a InsertarMaterial ??? -->r.HandleFunc("/material/nuevo", h.InsertarMaterial)
busca en la url y la funcion InsertarMaterial ...llama a nuevomaterial.html
que tiene un formulario que cuando se llena se convierte en POST enviando los datos
por medio de formvalue a el h.InsertarMaterial(handler o manejador...)este los recibe y si
estan bien convertidos con strconv los guarda en una funcion de DAL (Data Access Layer)
*/
func InsertarMaterial(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		u.RequestPagina(w, "nuevomaterial", nil)
		return
	}
	//cuando haces click en aceptar toma los datos, si no estan los InsertarMaterial, si esta Actualizar
	//si no es nil
	idMaterial := r.FormValue("idMaterial")
	descripcion := r.FormValue("descripcion")
	fecha := r.FormValue("fecha")
	precio, _ := strconv.ParseFloat(r.FormValue("precio"), 64)
	cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))
	dolar, _ := strconv.ParseFloat(r.FormValue("dolar"), 64)
	bhabilitado, _ := strconv.Atoi(r.FormValue("bhabilitado"))
	var err error
	if idMaterial == "" {
		// Lógica de inserción
		_, err = cd.InsertarMaterial(descripcion, precio, cantidad, fecha, dolar, bhabilitado)
	} else {
		// Lógica de actualización
		id, _ := strconv.Atoi(idMaterial)
		_, err = cd.Actualizar(descripcion, precio, cantidad, fecha, dolar, id)
	}

	// Lógica de redirección y manejo de errores común para ambos casos
	if err == nil {
		http.Redirect(w, r, "/materiales", http.StatusMovedPermanently)
		return
	}

	// Si hay un error, se renderiza la página con los datos
	material := m.Materiales{
		Descripcion: descripcion,
		Precio:      precio,
		Cantidad:    cantidad,
		Fecha:       fecha,
		Dolar:       dolar,
		Bhabilitado: 1,
	}
	u.RequestPagina(w, "insertar", material)
}
