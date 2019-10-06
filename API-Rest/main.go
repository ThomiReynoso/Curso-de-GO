package main

import (
		"fmt"
		"net/http"
		"log"  //libreria para capturar errores y mostrarlos por la consola
		"github.com/gorilla/mux"
		"encoding/json" //libreria para responder en json
		) 

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/peliculas", MovieList)
	router.HandleFunc("/peliculas/{id}", MovieShow)


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

func MovieList(w http.ResponseWriter, r *http.Request) {
		movies := Movies{
			//debe contener tantas instancias de Movie como desee que tenga el array/slice
			Movie{"Busqueda implacable", 2013, "Robin Hooper"},
			Movie{"Buscando a nemo", 2005, "Di caprio"},
			Movie{"Rapidos y furiosos", 2019, "Toretto"},
		}

		//fmt.Fprintf(w, "Listado de Peliculas")

		//devolvemos un json cuando se solicita acceso a esta URL
		json.NewEncoder(w).Encode(movies)

}

func MovieShow(w http.ResponseWriter, r *http.Request) {
		//con mux.Vars obtengo TODOS los param que le mando por URL
		//"params" va a ser un Array Asociativo
		params := mux.Vars(r)
		movie_id := params["id"]

		fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}
