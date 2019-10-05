package main

import (
		"fmt"
		"net/http"
		"log"  //libreria para capturar errores y mostrarlos por la consola
		) 

func main() {
	//Para crear una Ruta. le paso una func anonima (Callback)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		//Aca pongo lo que devuelve el server en esta ruta
		
		//Le paso el Writer, donde tiene que escribir 
		fmt.Fprintf(w, "Hola mundo desde mi nuevo servidor web con GO")
	})


	/*Metodo para que el servidor se levante en el puerto 8080 y 
	escuche peticiones*/ 
	server := http.ListenAndServe(":8080", nil)
	log.Fatal(server)


}

