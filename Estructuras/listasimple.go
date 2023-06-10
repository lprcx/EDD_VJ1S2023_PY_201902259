package estructuras

import "fmt"

type Lista struct {
	Inicio   *Nodo
	Longitud int
}

func (l *Lista) estaVacia() bool {
	return l.Longitud == 0
}

func (l *Lista) Insertar(numero int) {
	if l.estaVacia() {
		l.Inicio = &Nodo{valor: numero, siguiente: nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo{valor: numero, siguiente: nil}
		l.Longitud++
	}
}

func (l *Lista) Mostrar() {
	aux := l.Inicio

	for aux != nil {
		fmt.Println(aux.valor)
		aux = aux.siguiente
	}
}
