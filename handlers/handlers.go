package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	cd "github.com/angorita/loft/capaDatos"
	m "github.com/angorita/loft/models"
	u "github.com/angorita/loft/utilitarios"
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
		fmt.Println("Estoy en get ")
		oMateriales = cd.ListaMateriales()
	} else {
		if idMaterial == "" {
			oMateriales = cd.ListaMateriales()
			fmt.Println("idmaterial sin nada")
		} else {
			//entero convertido con atoi
			num, _ := strconv.Atoi(idMaterial)
			oMateriales = cd.FiltrarId(num)
			fmt.Println("combo filtrado..con Id ..", oMateriales)
		}
	}
	obj := ListaMaterialesForm{ListaMateriales: oMateriales, IdCombo: idMaterial}
	u.RequestPagina(w, "combo", obj)
}
func Principal(w http.ResponseWriter, r *http.Request) {
	u.RequestPagina(w, "index", nil)
}
func InsertarMaterial(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		u.RequestPagina(w, "insertar", nil)
		return
	} else {
		idMaterial := r.FormValue("idMaterial")
		descripcion := r.FormValue("descripcion")
		fecha := r.FormValue("fecha")
		precio, _ := strconv.ParseFloat(r.FormValue("precio"), 64)
		cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))
		dolar, _ := strconv.ParseFloat(r.FormValue("dolar"), 64)
		bhabilitado, _ := strconv.Atoi("bhabilitado")
		if idMaterial == "" {
			_, err := cd.InsertarMaterial(descripcion, precio, cantidad, fecha, dolar, bhabilitado)
			if err == nil {
				http.Redirect(w, r, "materiales", http.StatusMovedPermanently)

			} else {
				material := m.Materiales{Descripcion: descripcion,
					Precio: precio, Cantidad: cantidad, Fecha: fecha,
					Dolar: dolar, Bhabilitado: 1, ExisteError: true, MensajeError: err.Error()}
				u.RequestPagina(w, "insertar", material)
			}

		} else {
			id, _ := strconv.Atoi(idMaterial)
			_, err := cd.Actualizar(id, descripcion, precio, cantidad, fecha, dolar)
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
	fmt.Println("Parametro, su valor es: ", oMaterial)
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
