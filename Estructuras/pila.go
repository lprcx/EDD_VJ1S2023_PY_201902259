package estructuras

import "fmt"

type Pila struct {
	Primero  *NodoPila
	Longitud int
}

func (p *Pila) Push(Imagen string, cliente *Cliente) {
	if p.Longitud == 0 {
		nuevoNodo := &NodoPila{cliente: cliente, Imagen: Imagen, Siguiente: nil}
		p.Primero = nuevoNodo
		p.Longitud++
	} else {
		nuevoNodo := &NodoPila{cliente: cliente, Imagen: Imagen, Siguiente: p.Primero}
		p.Primero = nuevoNodo
		p.Longitud++
	}
}

func (p *Pila) Pop() {
	if p.Longitud == 0 {
		fmt.Println("No hay elementos en la pila")
	} else {
		p.Primero = p.Primero.Siguiente
		p.Longitud--
	}
}

func (p *Pila) Graficar() {
	nombre_archivo := "./pila.dot"
	nombre_imagen := "pila.jpg"
	texto := "digraph pila{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record]"
	aux := p.Primero
	texto += "nodo0 [label=\""
	for i := 0; i < p.Longitud; i++ {
		texto = texto + "|(ID: " + aux.cliente.Id + ", Imagen: " + aux.Imagen + ")"
		aux = aux.Siguiente
	}
	texto += "\"]; \n}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
