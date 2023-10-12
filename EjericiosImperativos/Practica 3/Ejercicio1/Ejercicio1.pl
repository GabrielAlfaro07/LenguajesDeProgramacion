%Define una regla que toma una lista de numeros como parametros y
% retorna la suma de los numeros.

% Caso base: La suma de una lista vacía es 0.
sumlist([], 0).

% Regla recursiva: La suma de una lista no vacía es la suma del primer
% elemento de la lista (Cabeza) y la suma del resto de la lista (Cola).
sumlist([Cabeza|Cola], Suma) :-
    sumlist(Cola, RestoSuma),
    Suma is Cabeza + RestoSuma.
