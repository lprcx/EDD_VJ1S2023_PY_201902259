package estructuras

import (
	"fmt"
	"math/rand"
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

func (c *Cola) RecorrerCola(id string, nombre string) {
	if c.Longitud == 0 {
		fmt.Println("No hay clientes en la cola")
	} else {
		aux := c.Primero
		for i := 0; i < c.Longitud; i++ {
			var nuevoid int
			_ = nuevoid
			if id == "X" {
				x, err := strconv.Atoi(aux.cliente.Id)
				if err != nil {
					fmt.Println("Error", err)
					return
				}
				nuevoid = x
				nuevoid = rand.Intn(9999)
				id = strconv.Itoa(nuevoid)
			}
			fmt.Println("Nombre: ", aux.cliente.Nombre, " Carnet: ", aux.cliente.Id)
			aux = aux.Siguiente
		}
	}
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
