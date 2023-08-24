package main

import (
	"fmt"
	"strconv"
)

// ValidarNumero convierte el tamaño en un entero y verifica si es válido.
// Retorna el número y booleano que indica si es válido.
func validarNumero(tamaño string) (int, bool) {
	numero, err := strconv.Atoi(tamaño)
	if err != nil {
		fmt.Println("Error: El valor indicado no es un número.")
		return 0, false
	} else if numero%2 == 0 {
		fmt.Println("Error: El número no es impar.")
		return 0, false
	} else if numero < 0 {
		fmt.Println("Error: El número no es positivo.")
		return 0, false
	} else {
		return numero, true
	}
}

// ObtenerTamaño solicita y valida el tamaño de la fila central.
// Retorna el tamaño válido.
func obtenerTamaño() int {
	var tamaño string
	fmt.Print("Ingrese el tamaño de la fila central (impar positivo): ")
	fmt.Scanln(&tamaño)
	numero, err := validarNumero(tamaño)
	if !err {
		panic(err)
	}
	return int(numero)
}

// PatronSuperior imprime la parte superior del diamante.
func patronSuperior(tamaño int) {
	for i := 0; i < tamaño/2+1; i++ {
		// Imprimir espacios
		for j := 0; j < tamaño/2-i; j++ {
			fmt.Print(" ")
		}
		// Imprimir asteriscos
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// PatronInferior imprime la parte inferior del diamante.
func patronInferior(tamaño int) {
	for i := tamaño/2-1; i >= 0; i-- {
		// Imprimir espacios
		for j := 0; j < tamaño/2-i; j++ {
			fmt.Print(" ")
		}
		// Imprimir asteriscos
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	tamaño := obtenerTamaño()
	patronSuperior(tamaño)
	patronInferior(tamaño)
}