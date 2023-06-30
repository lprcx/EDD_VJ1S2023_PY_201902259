package estructuras

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type BlockChain struct {
	Inicio        *NodoBloque
	BloqueCreados int
}

func (b *BlockChain) InsertarBloque(fecha string, biller string, customer string, payment string) {
	if b.BloqueCreados == 0 {
		cadenaFuncion := strconv.Itoa(b.BloqueCreados) + fecha + biller + customer + payment
		hash := SHA256(cadenaFuncion)
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.BloqueCreados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previousHash": "0000",
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque}
		b.Inicio = nuevoBloque
		b.BloqueCreados++
	} else {
		cadenaFuncion := strconv.Itoa(b.BloqueCreados) + fecha + biller + customer + payment
		hash := SHA256(cadenaFuncion)
		aux := b.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.BloqueCreados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previousHash": aux.Bloque["hash"],
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque, Anterior: aux}
		aux.Siguiente = nuevoBloque
		b.BloqueCreados++
	}
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hash := h.Sum(nil)
	hexaString = hex.EncodeToString(hash)
	return hexaString
}

func (b *BlockChain) ArregloFacturas() []RespuestaBloque {
	aux := b.Inicio
	var finalArreglo []RespuestaBloque
	for aux != nil {
		finalArreglo = append(finalArreglo, RespuestaBloque{Id: aux.Bloque["customer"], Factura: aux.Bloque["hash"]})
		aux = aux.Siguiente
		//falta validar que sea el empleado loggeado
	}
	return finalArreglo
}
