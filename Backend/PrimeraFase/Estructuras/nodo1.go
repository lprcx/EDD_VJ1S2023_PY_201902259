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
	idcliente  string /*Representacion de Cliente*/
	Imagen     string /*Representacion de Imagen*/
	idempleado string
	Siguiente  *NodoPila
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

type NodoArbol struct {
	Izquierdo         *NodoArbol
	Derecho           *NodoArbol
	Valor             int
	Altura            int
	Factor_Equilibrio int
}

type NodoMatrizDeAdyacencia struct {
	Siguiente *NodoMatrizDeAdyacencia
	Abajo     *NodoMatrizDeAdyacencia
	Valor     string
}

type EnvioMatriz struct {
	Padre   string
	Cliente string
	Imagen  string
	Filtros string
}

type NodoHash struct {
	Llave      int // -1
	Id_Cliente string
	Id_Factura string
}

type NodoBloque struct {
	Bloque    map[string]string
	Siguiente *NodoBloque
	Anterior  *NodoBloque
}

type NodoBloquePeticion struct {
	Timestamp string
	Biller    string
	Customer  string
	Payment   string
}

type RespuestaBloque struct {
	Id      string
	Factura string
}
