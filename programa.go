package main

import (
			"fmt"
	//		"os"
	//		"strconv"
			"time"
	   )  

func main () {
	
/*	
	fmt.Println("\n***** CONDICIONALES *******\n")

	//fmt.Println(os.Args) muestra dir del ejecutable
	
	fmt.Println("Hola, " + os.Args[1] + " bienvenido a mi programa GO" )
	
	//el "_" es una var que alojara el "error" que devuelve os.Args[2], lo obtengo y lo "desecho"
	edad,_ := strconv.Atoi(os.Args[2])
	//edad := 66
	
	if edad > 18 && edad < 60 {
		fmt.Println("Se encuentra entre 18 y 60, sos mayor de edad")
	} else if edad >= 60{
		fmt.Println("Sos anciano")
	}else{
		fmt.Println("Sos Menor de edad")
	} 

*/
	fmt.Println("\n***** BUCLES *******\n")
//For
/*	
	tope := 10
	for i := 0; i < tope; i++ {
		if i%2 == 0 {
			fmt.Println("El n°" + strconv.Itoa(i) + " Es par")
		}else{
			fmt.Println("El n°" + strconv.Itoa(i) + " Es impar")
		}
		
	}
*/
	
//	FOREACH
/*
	peliculas := []string{"Peli 1", "Peli 2", "peli 3", "peli 4"}

	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}
*/

//	SWITCH

	momento := time.Now()
	//el metodo Weekday se alimenta de "Now"
	hoy := momento.Weekday()

	switch hoy{
	case 0:
		fmt.Println("Hoy es Domingo")
	case 1:
		fmt.Println("Hoy es Lunes")
	case 2:
		fmt.Println("Hoy es Martes")
	case 5: 
		fmt.Println("Hoy es Viernes")	
	default: 
		fmt.Println("Es otro dia")		
	}
}