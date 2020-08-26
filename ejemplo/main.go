package main

import (
	"ejemplo/bd"
	"ejemplo/bienhechores"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//La funcion main es la primera funcion en ser llamada
	api()
}

func api() {
	//Se hace llamar a la funcion para crear una nueva conexion a la base de datos
	bd.NuevaConexionBD()
	//Se termina la conexion con la base de datos
	//el metodo 'defer' hará que esta función no sea ejecutada hasta que todos los demás procesos de la función 'api' se terminen
	defer bd.TerminarConcexionBD()

	//se crea el nuevo router para crear la api.
	//El metodo 'StrictSlash' cuando está en 'true' no lee la última diagonal '/' de la URL en caso de que la URL termine en diagonal
	//es decir que si se solicita '/get-membresia/' es lo mismo que solicitar '/get-membresia' en el navegador
	gorillaRoute := mux.NewRouter().StrictSlash(true)

	//Empezamos a crear nuestro arbol de direcciones con el metodo 'HandleFunc'
	//se coloca la URL que se estará escuchando, posteriormente la funcion que se ejecutará y por ultimo los metodos que permite este llamado
	gorillaRoute.HandleFunc("/api/get-membresia", bienhechores.GetMembresiaEndPoint).Methods("GET")
	gorillaRoute.HandleFunc("/api/post-membresia", bienhechores.PostMembresiaEndPoint).Methods("POST")

	//Con el metodo 'PathPrefix' estamos indicando apartir de que direccion se va a escuchar para publicar archivos
	//'http.FileServer' indica que va a estar imprimiendo en el navegaro archivos
	//'http.Dir' indica que todo lo que esté dentro de ese directorio se podrá ejecutar con 'http.FileServer' y dentro se indica el nombre de la carpeta
	gorillaRoute.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	//indicamos que toda la configuracion Handle y todo evento registra en la URL será manejado por la variable 'gorillaRoute'
	http.Handle("/", gorillaRoute)

	//Se lanza el servidor en el puerto :9090 y en caso de que haya algún error 'log.Fatal' nos lo hará saber
	log.Fatal(http.ListenAndServe(":9090", gorillaRoute))
}
