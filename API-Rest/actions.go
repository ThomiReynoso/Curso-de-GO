package main

import (
		"fmt"
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json" //libreria para responder en json
		) 

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
