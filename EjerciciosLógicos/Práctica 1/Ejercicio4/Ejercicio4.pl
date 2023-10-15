% Define un predicado que divide una lista con respecto a un elemento,
% llamado umbral, dado, que deja los valores menores a la izquierda y
% los valores mayores a la derecha.

% Caso base: Cuando la lista original es vacía, las listas resultantes también son vacías.
partir([], _, [], []).

% Caso 2: Si el primer elemento de la lista es menor o igual al umbral, se agrega a la lista de Menores
% y se llama a la recursión con el resto de la lista.
partir([X|Resto], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Resto, Umbral, Menores, Mayores).

% Caso 3: Si el primer elemento de la lista es mayor al umbral, se
% agrega a la lista de Mayores y se llama a la recursión con el resto
% de la lista.
partir([X|Resto], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Resto, Umbral, Menores, Mayores).
