window.onload = function () {
    Paginar("table")
}
function mostrarAlerta() {
    confirmacion().then((result) => {
        if (result.isConfirmed) {
            alerta()
        }
    });
}