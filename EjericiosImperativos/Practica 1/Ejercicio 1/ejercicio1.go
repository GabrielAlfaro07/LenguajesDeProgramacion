package main

import (
	"fmt"
	"os"
	"strings"
)

// Validar verifica si se produjo un error y provoca un pánico de ser el caso.
func validar(e error) {
	if e != nil {
		panic(e)
	}
}

// Estadisticas calcula y muestra la cantidad de líneas, palabras y caracteres.
func estadisticas(txt string) {
	texto := string(txt)
	lineas := len(strings.Split(texto, "\n"))
	palabras := len(strings.Fields(texto))
	caracteres := len(texto)

	fmt.Println("Cantidad de líneas:", lineas)
	fmt.Println("Cantidad de palabras:", palabras)
	fmt.Println("Cantidad de caracteres:", caracteres)
}

func main() {
	// Leer el contenido del archivo
	nombre := "world-history.txt"
	contenido, err := os.ReadFile(nombre)
	validar(err)

	// Calcular y mostrar las estadísticas
	fmt.Println("Nombre del archivo:", nombre)
	estadisticas(string(contenido))
}