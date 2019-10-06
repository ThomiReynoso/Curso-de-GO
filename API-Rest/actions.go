package main

import (
		"fmt"
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json" //libreria para responder en json
		"gopkg.in/mgo.v2"	
//		"gopkg.in/mgo.v2/bson"
		) 

//Creo un ARRAY GLOBAL 
var movies = Movies{
	//debe contener tantas instancias de Movie como desee que tenga el array/slice
	Movie{"Busqueda implacable", 2013, "Robin Hooper"},
	Movie{"Buscando a nemo", 2005, "Di caprio"},
	Movie{"Rapidos y furiosos", 2019, "Toretto"},
}	

//creo esta var global para reutilizarla cada vez que necesite acceder a la bd
var collection = getSession().DB("curso_go").C("movies")

//Devuelve un objeto de la LIBRERIA mgo.session
func getSession() *mgo.Session {

	//Estblezco Conexion con BD Mongo. En vez de localhost puedo poner la ip (127.0.0.1)
	session, err := mgo.Dial("mongodb://localhost")

	//'nil' es como 'nada'
	if err != nil{
		panic(err)
	}

	return session
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
	//recibo el json que llega en el Body de la Request
	decoder := json.NewDecoder(r.Body)

	//convierto el json a un objeto que pueda usar
	var movie_data Movie
	err := decoder.Decode(&movie_data) //decodifica el json y lo aloja en movie_data

	if err != nil {
		panic(err) //panic corta la ejecucion y muestra el error
	}

	//se usa para cerrar/limpiar la funcionalidad de algo
	defer r.Body.Close()



	//Inserto en la BD, en la coleccion Movies, la var movie_data (la cual contiene lo decodificado del json)
	/*comento sentencia debajo porque es mas optimo hacer una UNICA var global en vez de esto
		session := getSession()
		session.DB("curso_go").C("movies").Insert(movie_data)
	*/

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500) //Fallo
		return
	}
	//convierto a json el movie_data
	json.NewEncoder(w).Encode(movie_data)
	//escribo una respuesta Http
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) //Para escribir una Respuesta desde la api. 200 es que funciono ok


}

