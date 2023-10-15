// Definición del grafo como un mapa de conexiones entre nodos con sus respectivos costos.
let grafo = Map([("i", [("a", 2); ("b", 1)]);
                 ("a", [("i", 4); ("c", 2); ("d", 1)]);
                 ("b", [("i", 5); ("c", 3); ("d", 3)]);
                 ("c", [("a", 4); ("b", 6); ("x", 2)]);
                 ("d", [("a", 6); ("b", 5); ("f", 1)]);
                 ("f", [("d", 2)]);
                 ("x", [("c", 3)])])

// Definición del tipo Dijkstra para representar los nodos y costos en el algoritmo.
[<CustomEquality;CustomComparison>]
type Dijkstra<'N,'G when 'G:comparison>={toN:'N;cost:Option<'G>;fromN:'N}
                                        override g.Equals n = match n with
                                                           | :? Dijkstra<'N,'G> as n -> n.cost = g.cost
                                                           | _ -> false
                                        override g.GetHashCode() = hash g.cost
                                        interface System.IComparable with
                                          member n.CompareTo g =
                                            match g with
                                            | :? Dijkstra<'N,'G> as n when n.cost=None -> (-1)
                                            | :? Dijkstra<'N,'G> when n.cost=None -> 1
                                            | :? Dijkstra<'N,'G> as g -> compare n.cost g.cost
                                            | _ -> invalidArg "n" "expecting type Dijkstra<'N,'G>"

// Función Dijkstra para encontrar caminos más cortos en el grafo.
let inline Dijkstra N G y =
  let rec fN l f =
    if List.isEmpty l then f
    else
        let n = List.min l
        if n.cost = None then f
        else
            fN (l
                |> List.choose(fun n' ->
                    if n'.toN = n.toN then None
                    else
                        match n.cost, n'.cost, Map.tryFind (n.toN, n'.toN) G with
                        | Some g, None, Some wg -> Some {toN = n'.toN; cost = Some(g + wg); fromN = n.toN}
                        | Some g, Some g', Some wg when g + wg < g' -> Some {toN = n'.toN; cost = Some(g + wg); fromN = n.toN}
                        | _ -> Some n'))
                ((n.fromN, n.toN) :: f)
  let r = fN (N |> List.map (fun n -> {toN = n; cost = (Map.tryFind(y, n) G); fromN = y})) []
  (fun n ->
    let rec fN z l =
        match List.tryFind(fun (_, g) -> g = z) r with
        | Some(n', g') when y = n' -> Some(n' :: g' :: l)
        | Some(n', g') -> fN n' (g' :: l)
        | _ -> None
    fN n [])

// Definición de los nodos del grafo.
type Nodo = | A | B | C | D | E | F

// Definición de las conexiones y costos entre nodos.
let G = Map [((A, B), 7); ((A, C), 9); ((A, F), 14); ((B, C), 10); ((B, D), 15); ((C, D), 11); ((C, F), 2); ((D, E), 6); ((E, F), 9)]

// Ejemplo de uso del algoritmo de Dijkstra para encontrar rutas más cortas.
let rutas = Dijkstra [B; C; D; E; F] G A
let rutaAE = rutas E
let rutaAF = rutas F

// Función para imprimir la ruta en un formato legible
let printRuta ruta destino =
    match ruta with
    | Some camino -> 
        let caminoString = camino |> List.map string |> String.concat " -> "
        printfn "Ruta más corta a %A: %s" destino caminoString
    | None -> printfn "No hay ruta a %A" destino

printRuta rutaAE E
printRuta rutaAF F
