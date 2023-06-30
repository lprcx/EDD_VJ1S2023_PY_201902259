package estructuras

import (
	"fmt"
	"os"
)

func CrearArchivo() {

	var _, err = os.Stat("Archivo.json")

	if os.IsNotExist(err) {
		var file, err = os.Create("Archivo.json")
		if err != nil {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo creado exitosamente", "Archivo.json")
}

func EscribirArchivo(contenido string) {
	var file, err = os.OpenFile("Archivo.json", os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}

	err = file.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}
