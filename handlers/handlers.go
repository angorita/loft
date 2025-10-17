package handlers

import (
	m "X/models"
)

type ListaMaterialesForm struct {
	ListaMateriales []m.Materiales
	Id              int
	Nombre          string
	IdCombo         string
}
type ListaRNForm struct {
	ListaRN []m.RN
	Id      int
}
