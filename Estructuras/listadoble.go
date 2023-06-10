package estructuras

import (
	"fmt"
	"strconv"
)

type ListaDoble struct {
	Inicio   *Nodo2
	Longitud int
}

func (l *ListaDoble) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaDoble) Insertar(id string, nombre string, cargo string, password string) {
	nuevoEmpleado := &Empleado{Id: id, Nombre: nombre, Cargo: cargo, Password: password}
	// 0x8494AB8 = {20200000, Jose}
	//nuevoAlumno := &Alumno{carnet, nombre}
	if l.estaVacia() {
		l.Inicio = &Nodo2{nuevoEmpleado, nil, nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo2{nuevoEmpleado, nil, aux}
		l.Longitud++
	}
}

func (l *ListaDoble) MostrarAscendente() {
	aux := l.Inicio
	for aux != nil {
		fmt.Print(aux.empleado.Id)
		fmt.Println(" --> ", aux.empleado.Nombre)
		aux = aux.siguiente
	}
}

func (l *ListaDoble) MostrarDescente() {
	aux := l.Inicio
	for aux.siguiente != nil {
		aux = aux.siguiente
	}
	/*Imprimir hacia atras*/
	// Inicio = {Alumno, 0x0000001, nil}
	for aux != nil {
		fmt.Print(aux.empleado.Id)
		fmt.Println(" --> ", aux.empleado.Nombre)
		aux = aux.anterior
	}
}

func (l *ListaDoble) Reporte() {
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.empleado.Nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
