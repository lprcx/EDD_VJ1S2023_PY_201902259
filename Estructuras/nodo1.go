package estructuras

/*type Nodo struct {
	valor     int
	siguiente *Nodo
}*/

type Nodo2 struct {
	empleado  *Empleado
	siguiente *Nodo2
	anterior  *Nodo2
}

type NodoLista struct {
	empleado  *Empleado
	siguiente *NodoLista
}

type NodoCircular struct {
	cliente   *Cliente
	siguiente *NodoCircular
}

type NodoCola struct {
	Empleado  *Empleado /* Representacion de Cliente */
	Siguiente *NodoCola
}

type NodoPila struct {
	Empleado  *Empleado /*Representacion de Cliente*/
	Color     string    /*Representacion de Imagen*/
	Siguiente *NodoPila
}
