package Facturas

import (
	"EDD_VJ1S2023_PY_201902259/estructuras/GenerarArchivos"
	"EDD_VJ1S2023_PY_201902259/estructuras/TablaHash"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type BlockChain struct {
	Inicio          *NodoBloque
	Bloques_Creados int
}

func (b *BlockChain) InsertarBloque(fecha string, biller string, customer string, payment string) {
	cadenaFuncion := strconv.Itoa(b.Bloques_Creados) + fecha + biller + customer + payment
	hash := SHA256(cadenaFuncion)
	if b.Bloques_Creados == 0 {
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": "0000",
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque}
		b.Inicio = nuevoBloque
	} else {
		aux := b.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": aux.Bloque["hash"],
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque, Anterior: aux}
		aux.Siguiente = nuevoBloque
	}
	b.Bloques_Creados++
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hash := h.Sum(nil)
	hexaString = hex.EncodeToString(hash)
	return hexaString
}

func (b *BlockChain) InsertarTabla(tabla *TablaHash.TablaHash, idEmpleado string) {
	aux := b.Inicio
	for aux != nil {
		if aux.Bloque["biller"] == idEmpleado {
			tabla.Insertar(aux.Bloque["customer"], aux.Bloque["hash"])
		}
		aux = aux.Siguiente
	}
}

func (b *BlockChain) Graficar() {
	nombre_archivo := "./Reporte/reportepagos.dot"
	nombre_imagen := "./Reporte/reportepagos.jpg"
	texto := "digraph reportepagos{\n"
	texto += "rankdir=UD;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := b.Inicio
	contador := 0
	for i := 0; i < b.Bloques_Creados; i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{TimeStamp: " + aux.Bloque["timestap"] + "\n" + ", Biller: " + aux.Bloque["biller"] + ", Customer: " + aux.Bloque["customer"] + ", PreviousHash:" + aux.Bloque["previoushash"] + "}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < b.Bloques_Creados-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	strconv.Itoa(contador)
	texto += "}"
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(texto, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}
