package bd

import "ejemplo/estructura"

func GetMembresia() (membresias []estructura.Membresia) {
	//se crean dos variables en las que se almacenan los valores de forma temporal antes de guardarlos en el slice
	var (
		idMembresia   byte
		tipoMembresia string
	)

	//Se ejecuta el query, como se están solicitando 'n' cantidad de datos se usa 'Query'
	//Si se necesitara obtener un solo dato (como en el caso de requirir un dato mediante el id) se usaria 'QueryRow'
	membresia, _ := db.Query("SELECT IdMembresia, TipoMembresia FROM Membresia")
	revisarError(err)
	for membresia.Next() {
		//se aconsiguen los datos de la BD
		err = membresia.Scan(
			&idMembresia,
			&tipoMembresia,
		)
		revisarError(err)
		//Se almacenan en el slice para ser regresados
		membresias = append(membresias, estructura.Membresia{
			IdMembresia:   idMembresia,
			TipoMembresia: tipoMembresia,
		})
	}

	//Se retorna el slice 'membresias'
	return
}

func PostMembresia(membresia estructura.Membresia) error {
	//Se ejecuta el query para insertar datos
	//Exec tambien sirve con delete, drop, alter, update, etc
	//El primer valor no nos interesa porque solo estamos buscando que los datos se ingresen correctamente o de lo contrario saber el error, por eso se omite
	//Ese valor omitido nos podria servir para obtener el nuevo Id que se le ha otogado con la funcion 'LastInserID' pero no es el caso
	_, err := db.Exec("INSERT INTO Membresia(TipoMembresia) VALUES (?)", membresia.TipoMembresia)

	//Retornamos el error para usarlo directamente en la página, de no existir error el dato retornado sera 'null'
	return err
}
