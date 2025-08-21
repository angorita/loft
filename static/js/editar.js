function mostrarAlerta() {
    var descripcion = document.getElementById("txtDescripcion").value;
    var precio = document.getElementById("txtPrecio").value;
    var cantidad = document.getElementById("txtCantidad").value;
    var fecha = document.getElementById("txtFecha").value;
    var dolar = document.getElementById("txtDolar").value;
    var id=document.getElementById("txtId").value;
    var exito = true;
    
    var contenido = "<ul>";
    if (descripcion == "") {
        contenido += "<ol class='alert alert-danger mt-2'>Debe ingresar la descripcion</ol>"
        exito = false
    }
    if (precio == "") {
        contenido += "<ol class='alert alert-danger mt-2'>Debe ingresar el precio</ol>"
        exito = false
    }
    if (cantidad == "") {
        contenido += "<ol class='alert alert-danger mt-2'>Debe ingresar la cantidad</ol>"
        exito = false
    }
    if (fecha == "") {
        contenido += "<ol class='alert alert-danger mt-2'>Debe ingresar la fecha</ol>"
        exito = false
    }
    if (dolar == "") {
        contenido += "<ol class='alert alert-danger mt-2'>Debe ingresar el precio del dolar hoy</ol>"
        exito = false
    }

    contenido += "</ul>";
    if (exito) {
        return confirmacion().then((result) => {
            if (result.isConfirmed) {
                alerta()
                document.getElementById("frmEnviar").submit();
            }
        })
    }
    document.getElementById("divErrores").innerHTML = contenido;
}