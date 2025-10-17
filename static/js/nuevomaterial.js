function mostrarAlerta() {
    const descripcion = document.getElementById("txtDescripcion").value.trim();
    const precio = document.getElementById("txtPrecio").value.trim();
    const cantidad = document.getElementById("txtCantidad").value.trim();
    const fecha = document.getElementById("txtFecha").value.trim();
    const dolar = document.getElementById("txtDolar").value.trim();
    const bhabilitado = document.getElementById("chkBhabilitado").checked;

    const errores = [];

    // Lógica de validación (la misma que ya tenías)
    if (descripcion === "") {
        errores.push("Debe ingresar la descripción.");
    }
    if (precio === "" || isNaN(precio) || parseFloat(precio) <= 0) {
        errores.push("Debe ingresar un precio válido (número mayor que 0).");
    }
    if (cantidad === "" || isNaN(cantidad) || parseInt(cantidad, 10) <= 0) {
        errores.push("Debe ingresar una cantidad válida (número entero mayor que 0).");
    }
    if (fecha === "") {
        errores.push("Debe ingresar la fecha.");
    }
    if (dolar === "" || isNaN(dolar) || parseFloat(dolar) <= 0) {
        errores.push("Debe ingresar un precio del dólar válido (número mayor que 0).");
    }

    const divErrores = document.getElementById("divErrores");

    if (errores.length > 0) {
        // Muestra los errores si existen
        const contenido = `<div class="alert alert-danger mt-2"><ul>${errores.map(err => `<li>${err}</li>`).join('')}</ul></div>`;
        divErrores.innerHTML = contenido;
    } else {
        // Si no hay errores, envía el formulario directamente.
        divErrores.innerHTML = ""; // Limpia errores previos
        document.getElementById("frmEnviar").submit();
    }
}