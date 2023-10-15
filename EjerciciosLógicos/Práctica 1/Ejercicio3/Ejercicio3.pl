% Define una regla que recibe como parametro una lista con multiples
% listas anidadadas dentro, y retorna una lista con todos los mismos
% elementos dentro de las listas anidadas de forma lineal, sin listas
% anidadas.

% Caso base: Si la lista es vacía, la lista resultante también es vacía.
aplanar([], []).

% Predicado auxiliar para comprobar si un término es una lista.
is_list(X) :-
    is_list(X, []).

is_list([], _).
is_list([_|Cola], Cola).

% Caso 1: Si el primer elemento de la lista es una lista, aplana esa lista y
% concatena el resultado con la aplanación del resto de la lista.
aplanar([X|Resto], L2) :-
    is_list(X),  % Comprueba si X es una lista
    aplanar(X, XPlana),  % Aplana X
    aplanar(Resto, RestoPlano),  % Aplana el resto de la lista
    append(XPlana, RestoPlano, L2).  % Concatena los resultados.

% Caso 2: Si el primer elemento de la lista no es una lista, simplemente
% concatena ese elemento con la aplanación del resto de la lista.
aplanar([X|Resto], [X|RestoPlano]) :-
    \+ is_list(X),  % X no es una lista
    aplanar(Resto, RestoPlano).
