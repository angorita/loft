// Esta funcion esta en generic.js
window.onload = function () {
    Paginar("table")
}
function mostrarAlerta(id) {
    document.getElementById("txtid").value = id;
    confirmacion().then((result) => {
        if (result.isConfirmed) {
            var frm = document.getElementById("frm")
            frm.action = "/materiales/eliminar/" + id
            frm.submit()
        }
    })
}
