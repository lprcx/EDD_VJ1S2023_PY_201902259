package estructuras

type NodoDoble struct {
	nimagen   *Imagenn
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
	cliente   *Cliente /*Representacion de Cliente*/
	Imagen    string   /*Representacion de Imagen*/
	Siguiente *NodoPila
}

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Color     string
}
