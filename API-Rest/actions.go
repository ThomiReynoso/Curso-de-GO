package main

import (
		"fmt"
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json" //libreria para responder en json
		"gopkg.in/mgo.v2"	
		"gopkg.in/mgo.v2/bson"
		"log"
		
		) 

func getSession() *mgo.Session {

	//Estblezco Conexion con BD Mongo. En vez de localhost puedo poner la ip (127.0.0.1)
	session, err := mgo.Dial("mongodb://localhost")

	//'nil' es como 'nada'
	if err != nil{
		panic(err)
	}

	return session
}

//Para devolver un unico dato Movie
func responseMovie(w http.ResponseWriter, status int, results Movie) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) //Para escribir una Respuesta desde la api. 200 es que funciono ok

	//devolvemos un json cuando se solicita acceso a esta URL
	json.NewEncoder(w).Encode(results)

}

//Para devolver arrays de Movie
func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) 
	json.NewEncoder(w).Encode(results)

}

//creo esta var global para reutilizarla cada vez que necesite acceder a la bd
var collection = getSession().DB("curso_go").C("movies")

//Devuelve un objeto de la LIBRERIA mgo.session

func Index(w http.ResponseWriter, r *http.Request) {
		//Aca pongo lo que devuelve el server en esta ruta
		
		//Le paso el Writer, donde tiene que escribir 
		fmt.Fprintf(w, "Hola mundo desde mi nuevo servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {

		var results []Movie
		/*para hacer el find en MONGO 
		  Bindeamos lo que devuelve la consulta en "results" 
		  como param puedo mandarle lo mismo que uso para buscar en mongo */
		err := collection.Find(nil).Sort("-_id").All(&results) 
		
		if err != nil {
			log.Fatal(err)
		}else{
			fmt.Println("Resultados en MovieList: " , results)
		}
		
		//devuelvo un array	
		responseMovies(w, 200, results )

}

func MovieShow(w http.ResponseWriter, r *http.Request) {
		//con mux.Vars obtengo TODOS los param que le mando por URL
		//"params" va a ser un Array Asociativo
		params := mux.Vars(r)
		movie_id := params["id"]

		//Verifico si el dato es un objeto hexa para que acepte el json
		if !bson.IsObjectIdHex(movie_id) {
			w.WriteHeader(404) //fallo
			return
		}

		//codifico a json con bison la var (viene como string)
		//convierto a object id
		oid := bson.ObjectIdHex(movie_id)

	//	fmt.Println("movie_id: ", movie_id)
	//	fmt.Println("oid: ", oid)

		results := Movie{}
		//bindeo lo que traigo de la base a "results"
		err := collection.FindId(oid).One(&results)

	//	fmt.Println("results: ", results)

		if err != nil {
			w.WriteHeader(404) //fallo
			return
		}

		//Aca necesito devolver un solo dato
		responseMovie(w, 200, results )

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
	//escribo una respuesta Http
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) //Para escribir una Respuesta desde la api. 200 es que funciono ok

	//convierto a json el movie_data
	/*ES IMPORTANTE PONERLA COMO ULTIMA INSTRUCCION
	  PARA QUE SE TERMINE DEVOLVIENDO UN JSON	*/
	responseMovie(w, 200, movie_data )


}


func MovieUpdate(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		movie_id := params["id"]

		fmt.Println("movie_id", movie_id)

		//Verifico si el dato es un objeto hexa para que acepte el json
		if !bson.IsObjectIdHex(movie_id) {
			w.WriteHeader(404) //fallo
			return
		}

		//codifico a json con bison la var (viene como string)
		//convierto a object id
		oid := bson.ObjectIdHex(movie_id)
		
		//obtengo el objeto json que llega por el body
		decoder := json.NewDecoder(r.Body)
		
		var movie_data Movie
		err := decoder.Decode(&movie_data)
		fmt.Println("&movie_data", &movie_data)

		if err != nil {
			//si puede decodificarlo...
			panic(err) //muestra por consola el error
			w.WriteHeader(500)
			return
		}

		fmt.Println("oid", oid)
	
		//Dejamos de leer lo que hay en Body y lo limpiamos
		defer r.Body.Close()

		//Obtengo el Documento (Modelo) del id que estoy queriendo actualizar
		document := bson.M{"_id": oid}

		fmt.Println("oid", oid)
	
		//NUEVO dato a impactar en la base, en formato json
		change := bson.M{"$set": movie_data}

		err = collection.Update(document, change)
		fmt.Println("err", err)


		if err != nil {
			w.WriteHeader(404) //fallo
			return
		}

		//Aca necesito devolver un solo dato
		responseMovie(w, 200, movie_data )

}
