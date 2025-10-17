package models

import (
	"errors"
)

type Materiales struct {
	Id           int
	Descripcion  string
	Precio       float64
	Cantidad     int
	Total        float64
	Fecha        string
	Dolar        float64
	Bhabilitado  int
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
	if len(descripcion) > 100 {
		return errors.
			New("la descripcion tiene que tener maximo 100 ")
	}
	return nil
}
