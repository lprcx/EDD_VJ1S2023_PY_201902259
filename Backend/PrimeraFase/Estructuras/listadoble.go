package estructuras

import (
	"fmt"
	"strconv"
)

type ListaDoble struct {
	Inicio   *NodoDoble
	Longitud int
}

func (l *ListaDoble) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaDoble) Insertar(nimagen string, capas int) {
	nuevoImagen := &Imagenn{Imagen: nimagen, Capas: capas}
	// 0x8494AB8 = {20200000, Jose}
	//nuevoAlumno := &Alumno{carnet, nombre}
	if l.estaVacia() {
		l.Inicio = &NodoDoble{nuevoImagen, nil, nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoDoble{nuevoImagen, nil, aux}
		l.Longitud++
	}
}

func (l *ListaDoble) Recorrer(imagen string, capas int) *Imagenn {
	if l.estaVacia() {
		return nil
	} else {
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			if imagen == aux.nimagen.Imagen && capas == aux.nimagen.Capas {
				return aux.nimagen
			}
			aux = aux.siguiente
		}
		return nil
	}
}

func (l *ListaDoble) Mostrar() {
	if l.estaVacia() {
	} else {
		contador := 1
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			fmt.Println(strconv.Itoa(contador), ". ", aux.nimagen.Imagen)
			aux = aux.siguiente
			contador++
		}

	}
}

func (l *ListaDoble) Validarimagen(opcion string) string {
	aux := l.Inicio
	contador := 1
	for i := 0; i < l.Longitud; i++ {
		if opcion == strconv.Itoa(contador) {
			return aux.nimagen.Imagen
		}
		contador++
		aux = aux.siguiente
	}
	return "Opccion no valida"
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
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.nimagen.Imagen + "\"];\n"
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
