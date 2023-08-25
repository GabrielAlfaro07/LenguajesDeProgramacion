package main

import (
	"fmt"
	"sort"
)

// Definición de la estructura Producto para almacenar información de los productos.
type Product struct {
	name  string
	stock int
	price int
}

// Definición de un tipo Products para manejar una lista de Productos.
type Products []Product

// Lista global de productos.
var productsList Products

// Existencia mínima, utilizada para tomar decisiones sobre productos.
const minimumExistence int = 10

// Función para buscar un producto en la lista de productos.
// Retorna el producto encontrado y un valor de error (0 si no hay error, números mayores si hay algún error).
func (p *Products) searchProduct(name string) (*Product, int) {
	err := 1
	for i := 0; i < len(*p); i++ {
		if (*p)[i].name == name {
			err = 0
			return &(*p)[i], err
		}
	}
	return nil, err
}

// Función para agregar un producto a la lista de productos.
// Actualiza la existencia y el precio si el producto ya existe.
func (p *Products) addProduct(name string, stock int, price int) {
	product, err := p.searchProduct(name)
	if err == 0 {
		product.stock = product.stock + stock
		product.price = price
		fmt.Println("Producto actualizado.")
	} else {
		*p = append(*p, Product{name: name, stock: stock, price: price})
		fmt.Println("Producto agregado.")
	}
	fmt.Println()
}

// Función para agregar múltiples productos a la lista.
func (p *Products) addManyProducts(products ...Product) {
	for _, i := range products {
		p.addProduct(i.name, i.stock, i.price)
	}
}

// Función para vender productos.
// Reduce el stock y elimina el producto si no hay existencia.
func (p *Products) sellProduct(name string, sells int) {
	var index int
	product, err := p.searchProduct(name)
	if err == 0 {
		for i := 0; i < sells; i++ {
			if product.stock == 0 {
				for i := 0; i < len(*p); i++ {
					if (*p)[i].name == product.name {
						index = i
					}
				}
				*p = append((*p)[:index], (*p)[index+1:]...)
				fmt.Println("No hay stock. Producto eliminado.")
				break
			} else {
				product.stock--
				fmt.Println("Producto vendido.")
			}
		}
		fmt.Println()
	}
}

// Función para actualizar el precio de un producto.
func (p *Products) updatePrice(name string, price int) {
	product, err := p.searchProduct(name)
	if err == 0 {
		product.price = price
		fmt.Println("Precio actualizado.")
		fmt.Println()
	}
}

// Función para obtener productos con existencia mínima.
func (p *Products) minimumExistenceProducts() Products {
	var minExProducts Products
	for i := 0; i < len(*p); i++ {
		if (*p)[i].stock <= minimumExistence {
			minExProducts = append(minExProducts, (*p)[i])
		}
	}
	return minExProducts
}

// Función para reponer productos con existencia mínima.
func (p *Products) restockMinimumExistenceProducts(minExProducts Products) {
	for _, min := range minExProducts {
		min.stock = minimumExistence - min.stock
		p.addProduct(min.name, min.stock, min.price)
	}
}

// Función para ordenar productos por nombre.
func (p *Products) sortProductsByName() {
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].name < (*p)[j].name
	})
}

// Función para ordenar productos por precio.
func (p *Products) sortProductsByPrice() {
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].price < (*p)[j].price
	})
}

// Función para agregar datos iniciales a la lista de productos.
func addData() {
	p1 := Product{"Arroz", 15, 2500}
	p2 := Product{"Frijoles", 4, 2000}
	p3 := Product{"Leche", 8, 1200}
	p4 := Product{"Huevos", 20, 900}
	p5 := Product{"Carne", 16, 3000}

	productsList.addManyProducts(p1, p2, p3, p4, p5)

	productsList.addProduct("Cafe", 12, 4500)
}

func main() {
	addData()
	fmt.Println(productsList)

	productsList.addManyProducts(
		Product{"Arroz", 15, 2500},
		Product{"Carne", 4, 3000},
		Product{"Huevos", 8, 1000},
	)
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Actualizando el precio del arroz...")
	productsList.updatePrice("Arroz", 3000)
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Vendiendo productos...")
	productsList.sellProduct("Arroz", 20)
	fmt.Println(productsList)
	fmt.Println()

	productsList.sellProduct("Arroz", 20)
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Existencias mínimas:")
	fmt.Println(productsList.minimumExistenceProducts())
	fmt.Println()

	fmt.Println("Reponiendo existencias mínimas")
	productsList.restockMinimumExistenceProducts(productsList.minimumExistenceProducts())
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Ordenar por nombre:")
	productsList.sortProductsByName()
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Ordenar por precio:")
	productsList.sortProductsByPrice()
	fmt.Println(productsList)
	fmt.Println()
}