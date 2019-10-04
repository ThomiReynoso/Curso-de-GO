package main

import "fmt"

//Defino Structs

type Futbolista struct {
	apyn string
	equipo string
	edad int
	derecho bool
}

//import "time"
func main() {

/*	var n1 int = 18
	n2 := 6
 	var suma int = 8 + 10
	var prueba float32 = 3.5
	var name string = "Soy Thomas Reynoso"
	fmt.Println("Hola Mundo " + name)
	fmt.Println("La suma es", suma)
	fmt.Println("La suma es", prueba)
	//time.Sleep(1)
*/

/*	 var carlitos = Futbolista{
	  	"Carlos Tevez",
	 	"Boke",
	 	32,
	 	true}

	 fmt.Println(carlitos.equipo)
	 fmt.Println(carlitos.derecho)	

	 if n1 % n2 == 0  {
	 	fmt.Println("Son divisibles")
	 } else {
	 	fmt.Println("No Son divisibles")
	 }
*/

	 //LLAMADO A FUNCIONES

//	 prueba(2, 3)
//	 fmt.Println(retorno())
//	 enviandoVariosParams("boke", " es", " lamitad", " +1" )

	 //CLOSURES
//    fmt.Println( guantes( 40, "$") )

	 //ARRAYS UNIDIMENSIONALES
	arrays()

	 
}
 

//Funcion con parametros
func prueba( param1 int, param2 float32){

	fmt.Print("Probando ", param1)
	fmt.Print(" hola ", param2)
}

func retorno() (string, int) {
//Puedo devolver varios params a la vez pero NO me deja hacer 2 retunrns y distinto momento	
	var1 := "Hola"
	var2 := 44 

//	fmt.Println(var2) de esta forma muestra 1? el "44"
	return var1, var2
}

//CLOSURES
func guantes ( cant float32, moneda string) (string, float32, string) {

	//Cuando ejecuto el closure, asegurarme de que devuelva el tipo dato que voy a devolver en la func padre
	precio := func () float32{
		return cant*3
	}

	return "El precio del prod:" ,precio(), moneda
}

//con los "..." le indico que estoy esperando 1, 0 o muchos params
func enviandoVariosParams( parametros ... string) {
	for _, parametro := range parametros{
		fmt.Print(parametro)
	}
	fmt.Println("-------------")
	fmt.Println(parametros[1])
}

//DEFINICION DE ARRAYS
func arrays( ) {

//Forma 1 de creacion
	var array1[3] string
	array1[0] = "Esto"
	array1[1] = "es"
	array1[2] = "boke"

	fmt.Println(array1)
	fmt.Println(array1[2])

//Forma 2 de creacion

	array2 := [3]string{
		"tengo",
		"mucho", 
		"sueño"}

	fmt.Println(array2)
	fmt.Println(array2[1])
}