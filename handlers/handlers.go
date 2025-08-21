package handlers

import (
	cd "LoftSQLite3/capaDatos"
	m "LoftSQLite3/models"
	u "LoftSQLite3/utilitarios"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ListaMaterialesForm struct {
	ListaMateriales []m.Materiales
	Id              int
	Nombre          string
	IdCombo         string
}

func Materiales(w http.ResponseWriter, r *http.Request) {
	var nombreMaterial string
	var oMateriales []m.Materiales
	var id int
	nombreMaterial = r.FormValue(`nombre`)
	if r.Method == "GET" {
		oMateriales = cd.ListaMateriales()
	} else {
		//Post(click button )
		oMateriales = cd.FiltrarMateriales(nombreMaterial)
	}
	obj := ListaMaterialesForm{ListaMateriales: oMateriales, Nombre: nombreMaterial, Id: id}
	u.RequestPagina(w, "materiales", obj)
}

func Combo(w http.ResponseWriter, r *http.Request) {
	var oMateriales []m.Materiales
	idMaterial := r.FormValue("idMaterial")
	if r.Method == "GET" {
		oMateriales = cd.ListaMateriales()
	} else {
		if idMaterial == "" {
			oMateriales = cd.ListaMateriales()

		} else {
			//entero convertido con atoi
			num, _ := strconv.Atoi(idMaterial)

			oMateriales = cd.FiltrarId(num)
		}
	}
	obj := ListaMaterialesForm{ListaMateriales: oMateriales, IdCombo: idMaterial}
	u.RequestPagina(w, "combo", obj)
}
func Principal(w http.ResponseWriter, r *http.Request) {
	u.RequestPagina(w, "index", nil)
}
func InsertarMaterial(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u.RequestPagina(w, "insertar", nil)
	} else {
		idMaterial := r.FormValue("idMaterial")
		descripcion := r.FormValue("descripcion")
		fecha := r.FormValue("fecha")
		Num := r.FormValue("precio")
		Num2 := r.FormValue("cantidad")
		Num3 := r.FormValue("dolar")
		precio, _ := strconv.ParseFloat(Num, 64)
		cantidad, _ := strconv.Atoi(Num2)
		dolar, _ := strconv.ParseFloat(Num3, 64)
		if idMaterial == "" {
			_, err := cd.InsertarMaterial(descripcion, precio, cantidad, fecha, dolar, true)
			if err == nil {
				http.Redirect(w, r, "materiales", http.StatusMovedPermanently)

			} else {
				material := m.Materiales{Descripcion: descripcion,
					Precio: precio, Cantidad: cantidad, Fecha: fecha,
					Dolar: dolar, ExisteError: true, MensajeError: err.Error()}
				u.RequestPagina(w, "insertar", material)
			}

		} else {
			num, _ := strconv.Atoi(idMaterial)
			_, err := cd.Actualizar(num, descripcion, precio, cantidad, fecha, dolar)
			if err == nil {
				http.Redirect(w, r, "materiales", http.StatusMovedPermanently)
			}
		}
	}
}
func EditarMaterial(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	num, _ := strconv.Atoi(id)
	oMaterial := cd.BuscarMaterialesPorId(num)
	u.RequestPagina(w, "editar", oMaterial)
}
func EliminarMaterial(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	num, _ := strconv.Atoi(id)
	_, errorMaterial := cd.Eliminar(num)
	if errorMaterial == nil {
		http.Redirect(w, r, "/materiales", http.StatusMovedPermanently)
	}
}
