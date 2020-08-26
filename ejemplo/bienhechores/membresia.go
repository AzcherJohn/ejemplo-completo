package bienhechores

import (
	"ejemplo/bd"
	"ejemplo/estructura"
	"encoding/json"
	"net/http"
)

func GetMembresiaEndPoint(w http.ResponseWriter, r *http.Request) {
	//Se codifica la informacion en lenguaje 'json' y se imprime en el navegador 'json.NewEncoder(w)'
	//la informacion que se codifica es el resultado de la funcion 'bd.GetMembresia()' el cual será un slice
	json.NewEncoder(w).Encode(bd.GetMembresia())
}

func PostMembresiaEndPoint(w http.ResponseWriter, r *http.Request) {
	var membresia estructura.Membresia
	//Se decodifica del lengiaje 'json' la información del cuerpo 'r.Body' del formulario que es enviado
	//Y se almacena en la variable 'membresia'
	json.NewDecoder(r.Body).Decode(&membresia)
	//Se manda la variable 'membresia' a la funcion 'bd.PostMembresia' y el error que regrese se manda al navegador
	json.NewEncoder(w).Encode(bd.PostMembresia(membresia))

}
