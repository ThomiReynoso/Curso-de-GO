package main

import (
		"net/http"
		"log"  //libreria para capturar errores y mostrarlos por la consola
		) 

func main() {
	
	//Se añade este metodo, para mejorar la arquitectura de la API
	router := NewRouter()

	/*Metodo para que el servidor se levante en el puerto 8080 y 
	escuche peticiones. Recibe el objeto "Router", el cual tendra TODAS las 
	config de router y las distintas rutas que establezca*/ 
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)


}

