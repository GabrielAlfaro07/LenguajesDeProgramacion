% Define un predicado que, a partir de una lista de cadenas de
% caracteres (strings), filtra aquellas que contengan una subcadena que
% el usuario indique.

% Predicado para filtrar cadenas que contienen la subcadena especificada.
sub_cadenas(Subcadena, Lista, Filtradas) :-
    sub_cadenas(Subcadena, Lista, Filtradas, []).

% Caso base: Cuando la lista de entrada está vacía, terminamos y retornamos el resultado.
sub_cadenas(_, [], Filtradas, Filtradas).

% Predicado auxiliar para agregar elementos a una lista.
append([], L, L).
append([X|L1], L2, [X|L3]) :- append(L1, L2, L3).


% Regla recursiva: Procesa la cabeza de la lista y busca la subcadena.
sub_cadenas(Subcadena, [Cabeza|Resto], Filtradas, ResultadoParcial) :-
    (   sub_string(Cabeza, _, _, _, Subcadena)
    ->  % Si la subcadena se encuentra en Cabeza, la incluimos en la lista de Filtradas.
        append(ResultadoParcial, [Cabeza], NuevoResultado),
        sub_cadenas(Subcadena, Resto, Filtradas, NuevoResultado)
    ;   % Si no se encuentra la subcadena en Cabeza, seguimos con el resto de la lista.
        sub_cadenas(Subcadena, Resto, Filtradas, ResultadoParcial)
    ).
