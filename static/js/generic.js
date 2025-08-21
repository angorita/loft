//boton de sweet alert que fue personalizado
function confirmacion(titulo = "Confirma accion ?", subtitulo = "Por favor haga click en Aceptar o Cancelar") {
  Swal.fire({
    title: titulo,
    text: subtitulo,
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: "Guardar"
  })
}
function alerta(titulo = "Insertado", mensaje = "Se guardo Rebien!!") {
  Swal.fire(
    titulo,
    mensaje,
    'success'
  )
}
//cuando llaman a Paginar le pasan el id de la table y llama a Datatable, paginandola
function Paginar(idtabla) {
  $('#' + idtabla).DataTable()
}
function Pintar(url, cabeceras, camposMostrar, divImprimir = "divTabla",
  mostrarEditar = false, mostrarEliminar = false, propiedadId = "") {
  {
    var contenido = ""
    contenido += "<table class='table container' >"
    fetch(url).then(
      res => res.json
    ).then(res => {
      contenido += "<thead class='table-dark container' >"
      contenido += "<tr>"
      for (var i = 0; i < cabeceras.length; i++) {
        contenido += "<th>" + cabeceras[i] + "</th>"
      }
      if (mostrarEditar == true || mostrarEliminar == true) {
        contenido += "<th>Acciones</th>"
      }
      contenido += "</tr>"
      contenido += "</thead>"
      contenido += "<tbody>"
      //Por registro
      var fila;
      var nombrePropiedadAMostrar
      console.log(fila)
      for (var i = 0; i < res.length; i++) {
        contenido += "<tr>";
        //Objecto 
        fila = res[i]
        for (var j = 0; j < camposMostrar.length; j++) {
          nombrePropiedadAMostrar = camposMostrar[j]
          contenido += "<td>"
          contenido += fila[nombrePropiedadAMostrar]
          contenido += "</td>"
        }
        if (mostrarEditar == true || mostrarEliminar == true) {
          contenido += "<td>"
          if (mostrarEditar == true) {
            contenido += `
                
                <a class="btn btn-primary" 
                onclick='Editar(${fila[propiedadId]})'>
                <svg width="1em" height="1em" viewBox="0 0 16 16" class="bi bi-box-arrow-in-down-right" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" d="M14.5 13a1.5 1.5 0 0 1-1.5 1.5H3A1.5 1.5 0 0 1 1.5 13V8a.5.5 0 0 1 1 0v5a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5H9a.5.5 0 0 1 0-1h4A1.5 1.5 0 0 1 14.5 3v10z"/>
                    <path fill-rule="evenodd" d="M4.5 10a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-1 0v4.5H5a.5.5 0 0 0-.5.5z"/>
                    <path fill-rule="evenodd" d="M10.354 10.354a.5.5 0 0 0 0-.708l-8-8a.5.5 0 1 0-.708.708l8 8a.5.5 0 0 0 .708 0z"/>
                  </svg>
                </a>
                `
          }
          if (mostrarEliminar == true) {
            contenido += `
                <a class=" btn btn-danger " 
                onclick="mostrarAlerta(${fila[propiedadId]})">
                <svg width="1em" height="1em" viewBox="0 0 16 16" class="bi bi-trash" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                    <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                    <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4L4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                  </svg>
            </a>
                `
          }
          contenido += "</td>"
        }
        contenido += "</tr>"
      }
      contenido += "</tbody>"
      contenido += "</table>"
      document.getElementById(divImprimir).innerHTML = contenido;
    }
    )
  }

}
