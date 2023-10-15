// Definimos una función genérica llamada rotarLista que toma una lista,
// la cantidad de rotaciones y la dirección como argumentos.
let rotarLista (lista: 'T list) (rotaciones: int) (direccion: string) =
    // Obtenemos la longitud de la lista.
    let len = List.length lista
    // Calculamos el número total de rotaciones necesario, teniendo en cuenta la longitud de la lista.
    let rotacionesTotales = rotaciones % len // Para manejar desplazamientos mayores que la longitud de la lista
    // Utilizamos una expresión de coincidencia (match) para manejar la dirección especificada.
    match direccion with
    | "derecha" ->
        // Si la dirección es "derecha", dividimos la lista en dos partes: la parte izquierda y la parte derecha.
        let parteIzquierda = List.take (len - rotacionesTotales) lista
        let parteDerecha = List.skip (len - rotacionesTotales) lista
        // Concatenamos las partes derecha e izquierda para obtener la lista rotada a la derecha.
        parteDerecha @ parteIzquierda
    | "izquierda" ->
        // Si la dirección es "izquierda", dividimos la lista en dos partes: la parte izquierda y la parte derecha.
        let parteIzquierda = List.take rotacionesTotales lista
        let parteDerecha = List.skip rotacionesTotales lista
        // Concatenamos las partes derecha e izquierda para obtener la lista rotada a la izquierda.
        parteDerecha @ parteIzquierda
    | _ -> failwith "La dirección debe ser 'izquierda' o 'derecha'"

// Ejemplos de uso con listas de números enteros.
let listaInt = [1; 2; 3; 4; 5; 6; 7; 8; 9; 10]
let ListaIntIzquierda = rotarLista listaInt 4 "izquierda"
let ListaIntDerecha = rotarLista listaInt 3 "derecha"

// Imprimimos los resultados.
printfn "Lista de números original: %A" listaInt
printfn "Lista desplazada a la izquierda: %A" ListaIntIzquierda
printfn "Lista desplazada a la derecha: %A" ListaIntDerecha

// Ejemplos de uso con listas de cadenas de texto.
let listaString = ["Mateo"; "Julian"; "Eva"; "Alicia"; "Erick"; "Irene"; "Luca"; "Roberto"; ]
let ListaStringIzquierda = rotarLista listaString 2 "izquierda"
let ListaStringDerecha = rotarLista listaString 5 "derecha"

// Imprimimos los resultados.
printfn "Lista de cadenas original: %A" listaString
printfn "Lista desplazada a la izquierda: %A" ListaStringIzquierda
printfn "Lista desplazada a la derecha: %A" ListaStringDerecha