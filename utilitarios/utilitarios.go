package utilitarios

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"
)

func GenerarURL(uri, host, protocolo string, mapa map[string]string) string {
	u, _ := url.Parse(uri)
	u.Host = host
	u.Scheme = protocolo
	mapaFuncion := u.Query()
	for key, value := range mapa {
		mapaFuncion.Add(key, value)
	}
	u.RawQuery = mapaFuncion.Encode()
	return u.String()
}
func Request(metodo, url string) string {
	r, e := http.NewRequest(metodo, url, nil)
	if e != nil {
		panic("Hubo quilombo con el request")
	}
	cliente := &http.Client{}
	rpta, e := cliente.Do(r)
	if e != nil {
		panic("Quilombo con el Cliente")
	}
	bytes, e := ioutil.ReadAll(rpta.Body)
	if e != nil {
		panic("Quilombo al Leer !!")
	}
	return string(bytes)
}

// utilitarios.RequestPagina(w, "categoria", oCategoria)
func RequestPagina(w http.ResponseWriter, nombrePagina string, estructura interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := templateTodos.ExecuteTemplate(w, nombrePagina, estructura)
	if err != nil {
		w.WriteHeader(500)
		templateError.Execute(w, nil)
		fmt.Println("estoy en requestpagina", nombrePagina)
	}
}

var mapa = template.FuncMap{}
var templateTodos = template.Must(template.New("T").Funcs(mapa).ParseGlob("html/**/*.html"))
var templateError = template.Must(template.ParseFiles("html/error/error.html"))
