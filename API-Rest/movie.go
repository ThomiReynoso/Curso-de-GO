/*al agregarle package main, voy a poder acceder al contenido
de este archivo desde los archivos que tambien contengan
este package*/
package main


type Movie struct {
	Name string 	`json:"name"` //de esta manera renombro la propiedad cuando es devuelto como json, ya que es una mala practica que tenga nombre en mayusc
	Year int		`json:"year"`
	Director string `json:"director"`
}

type Movies []Movie