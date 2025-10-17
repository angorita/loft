package models

type RN struct {
	ID    int    `json:"id"`
	Nom   string `json:"nom"`
	FN    string `json:"fn"`
	AP    string `json:"ap"`
	AM    string `json:"am"`
	Parto string `json:"parto"`
	Obs   string `json:"obs"`
	Neo   string `json:"neo"`
	OS    string `json:"os"`
}
type ListaRN []RN
