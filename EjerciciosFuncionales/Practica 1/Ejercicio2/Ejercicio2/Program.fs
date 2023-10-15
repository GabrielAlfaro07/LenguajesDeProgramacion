// Definimos una función llamada filtrarPorSubcadena que toma una subcadena 
// y una lista de cadenas como argumentos.
let filtrarPorSubcadena (subcadena: string) (lista: string list) =
    // Utilizamos el operador |> para aplicar una operación en cadena.
    // En este caso, estamos aplicando List.filter para filtrar las cadenas de la lista.
    lista |> List.filter (fun cadena -> cadena.Contains(subcadena))

// Ejemplo de uso
let listaDeCadenas = ["manzana"; "banana"; "pera"; "naranja"; "sandía"]
let subcadena = "ana"

// Llamamos a la función filtrarPorSubcadena para obtener las cadenas que contienen la subcadena especificada.
let cadenasFiltradas = filtrarPorSubcadena subcadena listaDeCadenas

// Imprimimos los resultados en la consola.
printfn "Lista original: %A" listaDeCadenas
printfn "Subcadena a filtrar: %s" subcadena
printfn "Cadenas que contienen la subcadena: %A" cadenasFiltradas
