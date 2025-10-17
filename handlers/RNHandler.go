package handlers

import (
	cd "X/capaDatos"
	m "X/models"
	u "X/utilitarios"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RN(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**************************************************************************")
	fmt.Println("Estoy usando el HandleFunc", cd.ListaRN())
	fmt.Println("**************************************************************************")

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
