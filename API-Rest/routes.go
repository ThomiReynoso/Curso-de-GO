//Dentro de esta clase se encuentra TODA la configuracion necesaria para trabajar con rutas
package main

import (
		"net/http"
		"github.com/gorilla/mux"
		) 
type Route struct{
	Name 		string
	Method 		string //metodo http
	Pattern 	string
	HandleFunc	http.HandlerFunc	//si lo ponemos aca, evitamos hacerlo con cada ruta
}

type Routes []Route

//metodo que DEVUELVE una CONFIGURACION DE RUTAS (IMPORTANTE DEFINIRLO ASI, Y QUE DEVUELVA ESE TIPO DE DATO RUTA)
func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)

	//Vamos iterando por cada ruta definida en el array y la vamos creando
	for _, route := range routes{
		
		router.Methods(route.Method). //recibe la Propiedad, es el metodo http ( GET en este caso )
				Name(route.Name).
			    Path(route.Pattern).
				Handler(route.HandleFunc)
	}

	return router

}

var routes = Routes{
	Route{
		"Index"	,
		"GET"	,
		"/" 	,
		Index	,
	},
	Route{
		"MovieList"	,
		"GET"		,
		"/peliculas",
		MovieList	,
	},
	Route{
		"MovieShow"			,
		"GET"				,
		"/pelicula/{id}" 	,
		MovieShow			,
	},
	Route{
		"MovieAdd"	,
		"POST"		,
		"/pelicula"	,
		MovieAdd	,
	},
	Route{
		"MovieUpdate"	,
		"PUT"		,
		"/pelicula{id}"	,
		MovieUpdate	,
	},


}