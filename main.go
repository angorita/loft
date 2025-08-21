package main

import (
	h "LoftSQLite3/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/materiales", h.Materiales)
	r.HandleFunc("/combo", h.Combo)
	r.HandleFunc("/index", h.Principal)
	r.HandleFunc("/insertar", h.InsertarMaterial)
	r.HandleFunc("/editar/{id}", h.EditarMaterial)
	r.HandleFunc("/eliminar/{id}", h.EliminarMaterial)
	r.HandleFunc("/json", h.Wilder)
	fmt.Println(`http://localhost:8000/index`)
	http.ListenAndServe(":8000", r)
}
