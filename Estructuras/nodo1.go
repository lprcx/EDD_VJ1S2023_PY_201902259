package estructuras

type NodoDoble struct {
	imagen    *Imagenn
	siguiente *NodoDoble
	anterior  *NodoDoble
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
	cliente   *Cliente /* Representacion de Cliente */
	Siguiente *NodoCola
}

type NodoPila struct {
	Empleado  *Empleado /*Representacion de Cliente*/
	Color     string    /*Representacion de Imagen*/
	Siguiente *NodoPila
}
