package main

import (
	"errors"
	"fmt"
)

// Definición de la estructura Calzado para almacenar información de calzado.
type Calzado struct {
	modelo string
	precio int
	talla  int
}

// Método para agregar el modelo a la estructura Calzado.
func (c *Calzado) agregarModelo(modelo string) {
	c.modelo = modelo
}

// Método para agregar el precio a la estructura Calzado con validación.
func (c *Calzado) agregarPrecio(precio int) error {
	if precio < 0 {
		return errors.New("el precio debe ser mayor a 0")
	}
	c.precio = precio
	return nil
}

// Método para agregar la talla a la estructura Calzado con validación.
func (c *Calzado) agregarTalla(talla int) error {
	if talla < 34 || talla > 44 {
		return errors.New("la talla debe estar entre el 34 y el 44")
	}
	c.talla = talla
	return nil
}

// Definición de la estructura Inventario para almacenar información de inventario.
type Inventario struct {
	calzado Calzado
	stock   int
}

// Método para agregar información de calzado a la estructura Inventario.
func (i *Inventario) agregarCalzado(calzado Calzado) {
	i.calzado = calzado
}

// Método para agregar el stock a la estructura Inventario con validación.
func (i *Inventario) agregarStock(stock int) error {
	if stock < 0 {
		return errors.New("la cantidad de stock debe ser positiva")
	}
	i.stock = stock
	return nil
}

// Definición de la estructura Tienda para gestionar inventario y ventas.
type Tienda struct {
	inventario []Inventario
}

// Método para agregar información de inventario a la estructura Tienda.
func (t *Tienda) agregarInventario(inventario Inventario) {
	t.inventario = append(t.inventario, inventario)
}

// Método para realizar la venta de calzado en la estructura Tienda.
func (t *Tienda) venderCalzado(modelo string, talla int) error {
	for i, inventario := range t.inventario {
		if inventario.calzado.modelo == modelo && inventario.calzado.talla == talla {
			if inventario.stock < 0 {
				return errors.New("no hay stock disponible del calzado solicitado")
			} else {
				t.inventario[i].stock--
				return nil
			}
		}
	}
	return errors.New("no se encontró el calzado")
}

// Método para imprimir el inventario de la tienda.
func (t *Tienda) imprimirInventario() {
	for _, inventario := range t.inventario {
		fmt.Println("Modelo:", inventario.calzado.modelo)
		fmt.Println("Precio:", inventario.calzado.precio)
		fmt.Println("Talla:", inventario.calzado.talla)
		fmt.Println("Stock:", inventario.stock)
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	// Datos de ejemplo para el calzado
	modelos := []string{"Adidas", "Puma", "Skechers", "Nike", "Fila", "Converse"}
	precios := []int{1900, 436880, -17000, 12345, 90000, 5000}
	tallas := []int{40, 39, 36, 11, 50, 41}
	stocks := []int{190, 60, 75, 81, 63, -20}

	tiendaCalzado := Tienda{}

	// Creación de objetos Calzado, Inventario y agregado a la tienda
	for i := 0; i < 6; i++ {
		calzado := Calzado{}

		calzado.agregarModelo(modelos[i])

		err := calzado.agregarPrecio(precios[i])
		if err != nil {
			calzado.precio = 6000
			fmt.Println("En la posición", i, "=")
			print(err.Error())
			fmt.Println()
			continue
		}

		err = calzado.agregarTalla(tallas[i])
		if err != nil {
			calzado.talla = 40
			fmt.Println("En la posición", i, "=")
			print(err.Error())
			fmt.Println()
			continue
		}

		inventario := Inventario{}

		inventario.agregarCalzado(calzado)

		err = inventario.agregarStock(stocks[i])
		if err != nil {
			inventario.stock = 50
			fmt.Println("En la posición", i, "=")
			print(err.Error())
			fmt.Println()
			continue
		}

		tiendaCalzado.agregarInventario(inventario)
	}

	// Imprimir inventario inicial
	tiendaCalzado.imprimirInventario()

	// Simulación de ventas
	for i := 0; i < 70; i++ {
		err := tiendaCalzado.venderCalzado("Puma", 39)
		if err != nil {
			fmt.Println("En la cantidad", i, "=")
			fmt.Println(err.Error())
		} else {
			fmt.Println("Se ha vendido el calzado")
		}
	}

	fmt.Println()

	// Intento de venta cuando no hay stock
	err := tiendaCalzado.venderCalzado("Puma", 39)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Se ha vendido el calzado")
	}

	fmt.Println()

	// Imprimir inventario final
	tiendaCalzado.imprimirInventario()
}