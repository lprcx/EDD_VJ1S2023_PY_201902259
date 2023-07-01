package ArbolAVL

import "EDD_VJ1S2023_PY_201902259/estructuras/Peticiones"

type NodoArbol struct {
	Izquierdo         *NodoArbol
	Derecho           *NodoArbol
	Valor             *Peticiones.Pedido
	Altura            int
	Factor_Equilibrio int
}
