package main

import (
	"fmt"
	"sort"
)

type Product struct {
	name  string
	stock int
	price int
}

type Products []Product

var productsList Products

const minimumExistence int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (p *Products) searchProduct(name string) (*Product, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	err := 1
	for i := 0; i < len(*p); i++ {
		if (*p)[i].name == name {
			err = 0
			return &(*p)[i], err
		}
	}
	return nil, err
}

//modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
//tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
//que el uso de la nueva función ameriten

func (p *Products) addProduct(name string, stock int, price int) {
	product, err := p.searchProduct(name)
	if err == 0 {
		product.stock = product.stock + stock
		product.price = price
		fmt.Println("Product updated.")
	} else {
		*p = append(*p, Product{name: name, stock: stock, price: price})
		fmt.Println("Product added.")
	}
	fmt.Println()
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (p *Products) addManyProducts(products ...Product) {
	for _, i := range products {
		p.addProduct(i.name, i.stock, i.price)
	}
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista

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
				fmt.Println("No stock. Product erased.")
				break
			} else {
				product.stock--
				fmt.Println("Product sold.")
			}
		}
		fmt.Println()
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
		//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	}
}

func (p *Products) updatePrice(name string, price int) {
	product, err := p.searchProduct(name)
	if err == 0 {
		product.price = price
		fmt.Println("Price updated.")
		fmt.Println()
	}
}

//haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE

func (p *Products) minimumExistenceProducts() Products {
	var minExProducts Products
	for i := 0; i < len(*p); i++ {
		if (*p)[i].stock <= minimumExistence {
			minExProducts = append(minExProducts, (*p)[i])
		}
	}
	// debe retornar una nueva lista con productos con existencia mínima
	return minExProducts
}

func (p *Products) restockMinimumExistenceProducts(minExProducts Products) {
	for _, min := range minExProducts {
		min.stock = minimumExistence - min.stock
		p.addProduct(min.name, min.stock, min.price)
	}
}

func (p *Products) sortProductsByName() {
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].name < (*p)[j].name
	})
}

func (p *Products) sortProductsByPrice() {
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].price < (*p)[j].price
	})
}

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

	fmt.Println("Selling products...")
	productsList.sellProduct("Arroz", 20)
	fmt.Println(productsList)
	fmt.Println()

	productsList.sellProduct("Arroz", 20)
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Minimum existences:")
	fmt.Println(productsList.minimumExistenceProducts())
	fmt.Println()

	fmt.Println("Restocking minimum existences")
	productsList.restockMinimumExistenceProducts(productsList.minimumExistenceProducts())
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Sort by name:")
	productsList.sortProductsByName()
	fmt.Println(productsList)
	fmt.Println()

	fmt.Println("Sort by price:")
	productsList.sortProductsByPrice()
	fmt.Println(productsList)
	fmt.Println()
}
