package main

import (
	h "X/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Desde aqui se pide en material nuevo acceder al manejador multiple InsertarMaterial
Es un servidor muy sencillo, que  maneja las solicitudes que llegan a diferentes direcciones web (URLs) y responder a ellas de una manera específica.r := mux.NewRouter()
Aquí se crea un enrutador (o "router") llamado r. Este objeto es como el "cerebro" del servidor; su trabajo es decidir qué función ejecutar para cada URL que se visita.
r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
Esta línea es para servir archivos estáticos. Esto significa que si alguien visita una URL que empieza con /static/ (por ejemplo, http://localhost:8000/static/imagen.jpg), el servidor buscará un archivo llamado imagen.jpg dentro de una carpeta llamada static en tu proyecto. Es la forma de mostrar imágenes, hojas de estilo (CSS) y archivos JavaScript.
r.HandleFunc("/material/nuevo", h.InsertarMaterial)
r.HandleFunc("/materiales", h.Materiales)
r.HandleFunc("/combo", h.Combo)
r.HandleFunc("/index", h.Principal)
r.HandleFunc("/json", h.Wilder)
Estas líneas son las más importantes. Cada una asocia una URL con una función específica.

Si alguien visita la URL /material/nuevo, se ejecutará la función h.InsertarMaterial.

Si visitan /materiales, se ejecuta h.Materiales.

Y así sucesivamente para las demás URLs. La letra h antes del nombre de la función probablemente se refiere a un objeto que contiene los "handlers" (manejadores), es decir, las funciones que procesan las peticiones.
r.HandleFunc("/materiales/editar/{id}", h.EditarMaterial)
r.HandleFunc("/eliminar/{id}", h.EliminarMaterial)

Estas URLs son un poco más especiales porque tienen una parte variable, {id}. Esto le dice al enrutador que el valor de {id} puede cambiar. Por ejemplo, si alguien visita /materiales/editar/15, la función h.EditarMaterial se ejecutará y recibirá el número 15 como parte de la solicitud, sabiendo que debe editar el material con ese identificador.
http.ListenAndServe(":8000", r)
Esta es la línea que inicia el servidor. Le dice a la computadora: "Escucha las solicitudes que lleguen al puerto 8000 y usa el enrutador r para manejarlas". Si esta línea no se ejecuta, el servidor no funcionará.

*/

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//x rn=reciennacido
	r.HandleFunc("/rn", h.RN)
	//materiales
	r.HandleFunc("/material/nuevo", h.InsertarMaterial)
	r.HandleFunc("/materiales", h.Materiales)
	r.HandleFunc("/combo", h.Combo)
	r.HandleFunc("/index", h.Principal)
	r.HandleFunc("/materiales/editar/{id}", h.EditarMaterial)
	r.HandleFunc("/materiales/eliminar/{id}", h.EliminarMaterial) //eliminacion logica
	fmt.Println(`http://localhost:8000/index`)
	http.ListenAndServe(":8000", r)
}
