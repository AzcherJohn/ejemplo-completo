package estructura

type Membresia struct {
	//el nombre que se utilice en 'json' debe ser el mismo que se utilice en el front para enviar la informacion
	IdMembresia   byte   `json:"idMembresia"`
	TipoMembresia string `json:"tipoMembresia"`
}
