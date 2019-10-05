package main

import (
		"fmt"
		"net/http"
		"log"  //libreria para capturar errores y mostrarlos por la consola
		"github.com/gorilla/mux"
		) 

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)


	/*Metodo para que el servidor se levante en el puerto 8080 y 
	escuche peticiones. Recibe el objeto "Router", el cual tendra TODAS las 
	config de router y las distintas rutas que establezca*/ 
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)


}

func Index(w http.ResponseWriter, r *http.Request) {
		//Aca pongo lo que devuelve el server en esta ruta
		
		//Le paso el Writer, donde tiene que escribir 
		fmt.Fprintf(w, "Hola mundo desde mi nuevo servidor web con GO")
}

func Contact(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenido a la pagina de Contactos")
}
