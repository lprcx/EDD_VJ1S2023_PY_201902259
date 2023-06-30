package estructuras

import (
	"fmt"
	"strconv"
)

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

func (l *ListaEmpleados) Recorrer(id string, password string) *Empleado {
	if l.estaVacia() {
		return nil
	} else {
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			if id == aux.empleado.Id && password == aux.empleado.Password {
				return aux.empleado
			}
			aux = aux.siguiente
		}
		return nil
	}
}

func (l *ListaEmpleados) ReporteEmpleados() {
	nombreArchivo := "./listasimple.dot"
	nombreImagen := "./listasimple.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	_ = contador
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "ID: " + aux.empleado.Id + "Nombre: " + aux.empleado.Nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
