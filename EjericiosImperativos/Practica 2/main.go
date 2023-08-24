package main

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

//modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
//tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
//que el uso de la nueva función ameriten

func (l *listaProductos) venderProducto(nombre string) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		(*l)[prod].cantidad = (*l)[prod].cantidad - 1
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
		//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	}
}

//haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	return nil
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	//venda productos
}
