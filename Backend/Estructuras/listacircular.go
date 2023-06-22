package estructuras

import (
	"fmt"
	"strconv"
)

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
		fmt.Println("Nombre: ", aux.cliente.Nombre, " Id: ", aux.cliente.Id)
		aux = aux.siguiente
	}
}

func (l *ListaCircular) RecorrerClientes(id string, nombre string) *Cliente {
	if l.Longitud == 0 {
		return nil
	} else {
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			if id == aux.cliente.Id {
				return aux.cliente
			}
			aux = aux.siguiente
		}
		return nil
	}
}

func (l *ListaCircular) Validar(id string) bool {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		if id == aux.cliente.Id {
			return true
		}
		aux = aux.siguiente
	}
	return false
}

func (l *ListaCircular) ReporteClientes() {
	nombreArchivo := "./listacircular.dot"
	nombreImagen := "./listacircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Inicio
	contador := 0
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "ID: " + aux.cliente.Id + "Nombre: " + aux.cliente.Nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
