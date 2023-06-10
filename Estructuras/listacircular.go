package estructuras

import "fmt"

type ListaCircular struct {
	Inicio   *nodoLista
	Longitud int
}

func (l *ListaCircular) Insertar(id string, nombre string, password string) {
	nuevoEmpleado := &Empleado{Id: id, Nombre: nombre, Password: password}
	if l.Longitud == 0 {
		l.Inicio = &nodoLista{empleado: nuevoEmpleado, siguiente: nil}
		l.Inicio.siguiente = l.Inicio
		l.Longitud++
	} else {
		if l.Longitud == 1 {
			l.Inicio.siguiente = &nodoLista{empleado: nuevoEmpleado, siguiente: l.Inicio}
			l.Longitud++
		} else {
			aux := l.Inicio
			for i := 0; i < l.Longitud-1; i++ {
				aux = aux.siguiente
			}
			aux.siguiente = &nodoLista{empleado: nuevoEmpleado, siguiente: l.Inicio}
			l.Longitud++
		}
	}
}

func (l *ListaCircular) Mostrar() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Println("Nombre: ", aux.empleado.Nombre, " Carnet: ", aux.empleado.Id, " Password: ", aux.empleado.Password)
		aux = aux.siguiente
	}
}
