package estructuras

import "fmt"

type ListaCircular struct {
	Inicio   *NodoCircular
	Longitud int
}

func (l *ListaCircular) Insertar(id string, nombre string) {
	nuevoCliente := &Cliente{Id: id, Nombre: nombre}
	if l.Longitud == 0 {
		l.Inicio = &NodoCircular{cliente: nuevoCliente, siguiente: nil}
		l.Inicio.siguiente = l.Inicio
		l.Longitud++
	} else {
		if l.Longitud == 1 {
			l.Inicio.siguiente = &NodoCircular{cliente: nuevoCliente, siguiente: l.Inicio}
			l.Longitud++
		} else {
			aux := l.Inicio
			for i := 0; i < l.Longitud-1; i++ {
				aux = aux.siguiente
			}
			aux.siguiente = &NodoCircular{cliente: nuevoCliente, siguiente: l.Inicio}
			l.Longitud++
		}
	}
}

func (l *ListaCircular) Mostrar() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Println("Nombre: ", aux.cliente.Nombre, " Carnet: ", aux.cliente.Id)
		aux = aux.siguiente
	}
}
