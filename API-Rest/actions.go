package main

import (
		"fmt"
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json" //libreria para responder en json
		"log"
		) 

//Creo un ARRAY GLOBAL 
var movies = Movies{
	//debe contener tantas instancias de Movie como desee que tenga el array/slice
	Movie{"Busqueda implacable", 2013, "Robin Hooper"},
	Movie{"Buscando a nemo", 2005, "Di caprio"},
	Movie{"Rapidos y furiosos", 2019, "Toretto"},
}	

func Index(w http.ResponseWriter, r *http.Request) {
		//Aca pongo lo que devuelve el server en esta ruta
		
		//Le paso el Writer, donde tiene que escribir 
		fmt.Fprintf(w, "Hola mundo desde mi nuevo servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {

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

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	//recibo y decodifico el json que llega en el Body de la Request
	decoder := json.NewDecoder(r.Body)

	//lo convierto a un objeto que pueda usar
	var movie_data Movie
	err := decoder.Decode(&movie_data) //decodifica y lo aloja en movie_data


	if err != nil {
		panic(err) //panic corta la ejecucion y muestra el error
	}

	//se usa para cerrar/limpiar la funcionalidad de algo
	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)

}

