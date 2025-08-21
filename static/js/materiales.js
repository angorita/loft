// Esta funcion esta en generic.js
window.onload = function () {
    Paginar("table")
}
function mostrarAlerta(id){
   document.getElementById("txtid").value=id;
    confirmacion().then((result)=>{
        if(result.value){
            //si entro
            // alert(id)
            var frm=document.getElementById("frm")
            frm.action="eliminar/"+id
            frm.submit()
        }
    })
}