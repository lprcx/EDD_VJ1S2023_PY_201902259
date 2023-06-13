package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

/*func (colaAux *Cola) LeerArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		colaAux.Encolar(linea[0], linea[1])
	}
}*/

/*func (l *ListaCircular) LeerArchivo(ruta string)*/
func (l *ListaEmpleados) LeerArchivo(ruta string) []Empleado {
	var lista []Empleado
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return nil
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		/*l.Insertar(linea[0],linea[1])*/
		lista = append(lista, Empleado{Id: linea[0], Nombre: linea[1]})
	}
	return lista
}
