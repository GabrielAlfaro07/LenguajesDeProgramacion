package main

import "fmt"

// RotarLista realiza rotaciones en una lista en la dirección especificada.
// Recibe un puntero a la lista, la cantidad de rotaciones y la dirección (0: izquierda, 1: derecha).
// Retorna un error en caso de que la longitud de una lista sea 0.
func rotarLista[lista any](elementos *[]lista, rotaciones int, direccion int) error {
	longitud := len(*elementos)
	if longitud == 0 {
		return fmt.Errorf("la longitud de la lista es igual a 0")
	}

	fmt.Println("Cantidad de rotaciones =", rotaciones)
	rotaciones = rotaciones % longitud

	if direccion == 0 {
		fmt.Println("Direccion = Izquierda")
		*elementos = append((*elementos)[rotaciones:], (*elementos)[:rotaciones]...)
	} else {
		fmt.Println("Direccion = Derecha")
		*elementos = append((*elementos)[longitud-rotaciones:], (*elementos)[:longitud-rotaciones]...)
	}
	return nil
}

func main() {
	// Declaración de listas de diferentes tipos
	listaInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listaString := []string{"Gabriel", "Irene", "Isabel", "Erick", "Eduardo", "Javier", "Mauricio", "Ana"}
	listaBool := []bool{true, true, false, false, true, false, true, false}

	// Obteniendo punteros a las listas
	punteroInt := &listaInt
	punteroString := &listaString
	punteroBool := &listaBool

	rotaciones := 3

	// Realizando rotaciones y mostrando resultados
	fmt.Println("Secuencia original:", listaInt)
	err1 := rotarLista(punteroInt, rotaciones, 0)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Secuencia final rotada:", listaInt)
	fmt.Println()

	fmt.Println("Secuencia original:", listaString)
	err2 := rotarLista(punteroString, rotaciones, 1)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Secuencia final rotada:", listaString)
	fmt.Println()

	fmt.Println("Secuencia original:", listaBool)
	err3 := rotarLista(punteroBool, rotaciones, 0)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("Secuencia final rotada:", listaBool)
	fmt.Println()
}