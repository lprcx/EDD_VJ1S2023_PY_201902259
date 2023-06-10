package main

import (
	//estructuras "EDD_VJ1S2023_PY_201902259/Estructuras"
	"fmt"
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
			fmt.Println("Iniciar Sesión")
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
	if id == "admin" && password == "admin" {
		fmt.Println("Bienvenido Administrador :)")
		//menu_admin()
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
