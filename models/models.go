package models

import (
	"errors"
)

type Materiales struct {
	Descripcion  string
	Precio       float64
	Cantidad     int
	Total        float64
	Fecha        string
	Dolar        float64
	Id           int
	Bhabilitado  bool
	ExisteError  bool
	MensajeError string
}
type Wilder struct {
	Descripcion string
	Precio      float64
	Cantidad    int
}
type ListaMateriales []Materiales
type ListaWilder []Wilder

func Max(descripcion string) error {
	if len(descripcion) > 150 {
		return errors.
			New("la descripcion tiene que tener maximo 100 ")
	}
	return nil
}
