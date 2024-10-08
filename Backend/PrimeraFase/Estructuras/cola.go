package estructuras

import (
	"fmt"
	"strconv"
)

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

func (c *Cola) Encolar(id string, nombre string) {
	nuevoCliente := &Cliente{Id: id, Nombre: nombre}
	if c.Longitud == 0 {
		nuevoNodo := &NodoCola{nuevoCliente, nil}
		c.Primero = nuevoNodo
		c.Longitud++
	} else {
		nuevoNodo := &NodoCola{nuevoCliente, nil}
		aux := c.Primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay clientes en la cola")
	} else {
		c.Primero = c.Primero.Siguiente
		c.Longitud--
	}
}

func (c *Cola) MostrarPrimero() {
	fmt.Println(c.Primero.cliente.Id, " ", c.Primero.cliente.Nombre, " ")
}

func (c *Cola) RecorrerCola(id string, nombre string) *Cliente {
	if c.Longitud == 0 {
		fmt.Println("No hay clientes en la cola")
		return nil
	} else {
		aux := c.Primero
		for i := 0; i < c.Longitud; i++ {
			if id == aux.cliente.Id && nombre == aux.cliente.Nombre {
				return aux.cliente
			}
			//fmt.Println("Nombre: ", aux.cliente.Nombre, " Carnet: ", aux.cliente.Id)
			aux = aux.Siguiente
		}
		return nil
	}
}

func (c *Cola) ObtenerId() string {
	aux := c.Primero
	if aux != nil {
		return aux.cliente.Id
	} else {
		return "Está vacío"
	}
}

func (c *Cola) ObtenerNombre() string {
	aux := c.Primero
	if aux != nil {
		return aux.cliente.Nombre
	} else {
		return "Está vacío"
	}
}

func (c *Cola) ObtenerLongitud() int {
	return c.Longitud
}

func (c *Cola) Enlistar() *Cliente {
	if c.Longitud == 0 {
		return nil
	} else {
		return c.Primero.cliente
	}
}

func (c *Cola) Graficar() {
	nombre_archivo := "./cola.dot"
	nombre_imagen := "cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.Primero
	contador := 0
	for i := 0; i < c.Longitud; i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{ID: " + aux.cliente.Id + ", Nombre: " + aux.cliente.Nombre + "|}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < c.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
