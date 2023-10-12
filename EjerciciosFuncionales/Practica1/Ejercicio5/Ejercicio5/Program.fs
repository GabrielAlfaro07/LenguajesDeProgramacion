let laberinto = [
    ("INICIO", ["2"]);
    ("1", ["7"]);
    ("2", ["INICIO"; "8"; "3"]);
    ("3", ["2"; "9"; "4"]);
    ("4", ["10"]);
    ("5", ["11"; "6"]);
    ("6", ["5"]);
    ("7", ["1"; "13"]);
    ("8", ["2"; "9"]);
    ("9", ["8"; "3"]);
    ("10", ["4"; "16"]);
    ("11", ["5"; "17"]);
    ("12", ["18"]);
    ("13", ["7"; "14"]);
    ("14", ["13"; "15"; "20"]);
    ("15", ["14"; "21"]);
    ("16", ["10"; "22"]);
    ("17", ["11"; "23"]);
    ("18", ["12"; "24"]);
    ("19", ["25"]);
    ("20", ["14"; "26"]);
    ("21", ["15"; "22"]);
    ("22", ["16"; "21"]);
    ("23", ["17"; "29"]);
    ("24", ["18"; "30"]);
    ("25", ["19"; "31"]);
    ("26", ["20"; "27"]);
    ("27", ["26"; "28"]);
    ("28", ["27"; "34"; "29"]);
    ("29", ["23"; "28"]);
    ("30", ["24"; "36"]);
    ("31", ["25"; "32"]);
    ("32", ["31"; "33"; "FIN"]);
    ("33", ["32"; "34"]);
    ("34", ["33"; "28"; "35"]);
    ("35", ["34"; "36"]);
    ("36", ["35"; "30"]);
    ("FIN", ["32"]);
]

// Función para verificar si un elemento está en una lista.
let miembro elem lista =
    List.exists (fun x -> x = elem) lista

// Función para obtener los vecinos de un nodo en el laberinto.
let rec obtenerVecinos nodo grafo =
    match grafo with
    | [] -> []
    | (cabeza, vecinos) :: resto ->
        if cabeza = nodo then vecinos
        else obtenerVecinos nodo resto

// Función para extender una ruta en el laberinto.
let rec extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map  (fun x -> if (miembro x ruta) then [] else x::ruta) (obtenerVecinos (List.head ruta) grafo)) 

// Función principal para buscar una ruta en profundidad en el laberinto.
let rec busquedaEnProfundidad ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)  ) grafo
    prof_aux fin [[ini]] grafo

// Función para obtener la ruta más corta de una lista de rutas.
let rec obtenerRutaMasCorta listaRutas =
    match listaRutas with
    | [] -> [] 
    | [listaUnica] -> listaUnica 
    | cabeza :: resto ->
        let rutaMasCortaResto = obtenerRutaMasCorta resto 
        if List.length cabeza <= List.length rutaMasCortaResto then cabeza
        else rutaMasCortaResto

// Búsqueda de ruta desde el nodo INICIO al nodo FIN en el laberinto.
let rutaEncontrada = busquedaEnProfundidad "INICIO" "FIN" laberinto
printfn "Rutas encontradas: %A" rutaEncontrada

// Obtención de la ruta más corta entre INICIO y FIN.
let rutaMasCorta = obtenerRutaMasCorta rutaEncontrada
printfn "Ruta más corta: %A" rutaMasCorta
