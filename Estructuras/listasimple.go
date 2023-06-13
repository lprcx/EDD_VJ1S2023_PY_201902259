package estructuras

import "fmt"

type ListaEmpleados struct {
	Inicio   *NodoLista
	Longitud int
}

func (l *ListaEmpleados) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaEmpleados) Insertar(id string, nombre string, cargo string, password string) {
	nuevoEmpleado := &Empleado{Id: id, Nombre: nombre, Cargo: cargo, Password: password}
	if l.estaVacia() {
		l.Inicio = &NodoLista{empleado: nuevoEmpleado, siguiente: nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoLista{empleado: nuevoEmpleado, siguiente: nil}
		l.Longitud++
	}
}

func (l *ListaEmpleados) Mostrar() {
	aux := l.Inicio

	for aux != nil {
		fmt.Println(aux.empleado.Id, aux.empleado.Nombre, aux.empleado.Cargo, aux.empleado.Password)
		aux = aux.siguiente
	}
}
