%Define una regla que obtiene como parametros dos listas S y S1, que
%retorna verdadero si S1 es subconjunto de S.

%Caso base: Un conjunto vac�o es subconjunto de cualquier conjunto.
subconj(_, []).

% Regla recursiva: S1 es un subconjunto de S si el primer elemento de S1
% est� en S y el resto de S1 tambi�n es un subconjunto de S.
subconj(S, [X|S1]) :-
    member(X, S),
    subconj(S, S1).
