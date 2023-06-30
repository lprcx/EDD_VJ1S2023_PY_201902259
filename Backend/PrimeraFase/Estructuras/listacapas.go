package estructuras

import (
	"fmt"
	"strconv"
)

type Capa struct {
	layer_capa string
	file_capa  string
}

type NodoC struct {
	data      *Capa
	siguiente *NodoC
}

type ListaCapa struct {
	Inicio   *NodoC
	Longitud int
}

func (lista *ListaCapa) estaVacia() bool {
	if lista.Longitud == 0 {
		return true
	}
	return false
}

// Inserta al final
func (lista *ListaCapa) InsertaCapa(layer string, file string) {
	capas := &Capa{layer_capa: layer, file_capa: file}
	if lista.estaVacia() {
		lista.Inicio = &NodoC{data: capas, siguiente: nil}
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.siguiente != nil { //esto simula un while
			aux = aux.siguiente
		}
		aux.siguiente = &NodoC{data: capas, siguiente: nil}
		lista.Longitud++
	}
}

func (lista *ListaCapa) ListarCapa() {
	aux := lista.Inicio
	var contador int = 1
	for aux != nil {
		fmt.Println(strconv.Itoa(contador)+". ", aux.data.file_capa, "capa: ", aux.data.layer_capa)
		aux = aux.siguiente
		contador++
	}
}

func (lista *ListaCapa) BuscarCapa(opcion string) string {
	aux := lista.Inicio
	var contador int = 1
	for aux != nil {
		if opcion == strconv.Itoa(contador) {
			return aux.data.file_capa
		}
		//fmt.Println(strconv.Itoa(contador)+". ", aux.data.name_Imagen)
		aux = aux.siguiente
		contador++
	}
	return "Elija una opcion valida"
}

func NewListaSimpleC() *ListaCapa {
	return &ListaCapa{nil, 0}
}
