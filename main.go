package main

import (
	//estructuras "EDD_VJ1S2023_PY_201902259/Estructuras"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

var (
	id       string
	nombre   string
	cargo    string
	password string
	opcion3  int
)

func main() {
	opcion1 := 0
	salir := false
	for !salir {
		fmt.Println("---------------Login---------------")
		fmt.Println("1. Iniciar Sesión")
		fmt.Println("2. Salir del Sistema")
		fmt.Println("Elige una opción: ")
		fmt.Scanln(&opcion1)
		if opcion1 == 1 {
			main2()
		} else {
			salir = true
		}
	}

}

func main2() {
	salir := false
	_ = salir
	fmt.Println("Ingresa tu usuario: ")
	fmt.Scanln(&id)
	fmt.Println("Ingresa tu password")
	fmt.Scanln(&password)
	if id == "ADMIN_201902259" && password == "Admin" {
		fmt.Println("Bienvenido Administrador :)")
		menu_admin()
	} else if id != "ADMIN_201902259" && password != "Admin" {
		//verificacion()
	} else {
		fmt.Println("Datos Erróneos")
	}

}

func menu_admin() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("-------------Dashboard Administrador 201902259-----------")
		fmt.Println("1. Cargar Empleados")
		fmt.Println("2. Cargar Imagenes")
		fmt.Println("3. Cargar Usuarios")
		fmt.Println("4. Actualizar Cola")
		fmt.Println("5. Reportes de Estrucutras")
		fmt.Println("6. Cerrar Sesión")
		fmt.Print("************************************")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var ruta string
			fmt.Println("Ingrese la ruta del archivo")
			fmt.Println("")
			fmt.Scanln(&ruta)
			LeerArchivo(ruta)
		case 2:

		case 3:

		case 4:

		case 5:

		case 6:
			fmt.Println("Cerrando Aplicacion....")
			main()

		default:
			fmt.Println("Opción inválida")
		}

	}

}

func LeerArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		return
	}
	defer file.Close()

	leer := csv.NewReader(file)
	leer.Comma = ','
	//leer.FieldsPerRecord = -1

	encabezado, err := leer.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("Se ha cargado el archivo", encabezado)

	for {
		linea, err := leer.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No se pudo leer la linea")
			break
		}
		//cola.Encolar(strings.TrimSpace(linea[1]), "", strings.TrimSpace(linea[0]), strings.TrimSpace(linea[2]))
		fmt.Println("nombre: ", linea[1]+" Cargo"+linea[2], " id: ", linea[0])
	}
}
