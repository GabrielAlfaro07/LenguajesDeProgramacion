// Función que asigna un valor específico en una posición dada en una lista.
let mapPosicion posicion lista =
    List.mapi (fun indice x -> if indice = posicion then x else -1) lista

// Función que filtra todos los valores diferentes a -1 en una lista.
let filtrarMenosUno lista =
    List.filter (fun x -> x <> -1) lista

// Función para obtener el elemento en la posición n de una lista.
let n_esimo posicion lista = 
    if posicion >= List.length lista then -1
    else
       let listaFiltrada = mapPosicion posicion lista // Asigna -1 en todas las posiciones excepto en la especificada.
       let resultadoFiltrado = filtrarMenosUno listaFiltrada // Filtra los valores diferentes a -1.
       resultadoFiltrado.[0] // Retorna el primer elemento (único) de la lista resultante.

// Lista de números.
let numeros = [1; 2; 3; 4; 5; 6; 7; 8; 9; 10]

// Llama a la función n_esimo con posición 0 en la lista de números.
let resultado1 = n_esimo 0 numeros
let resultado2 = n_esimo 5 numeros
let resultado3 = n_esimo 9 numeros
let resultado4 = n_esimo 15 numeros


// Imprime el resultado.
printfn "Resultado 1: %d" resultado1
printfn "Resultado 2: %d" resultado2
printfn "Resultado 3: %d" resultado3
printfn "Resultado 4: %d" resultado4
