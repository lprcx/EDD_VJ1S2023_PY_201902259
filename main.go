package main

import (
	estructuras "EDD_VJ1S2023_PY_201902259/Estructuras"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var ListaSimple = estructuras.ListaEmpleados{Inicio: nil, Longitud: 0}
var ListaCircular = estructuras.ListaCircular{Inicio: nil, Longitud: 0}
var ListaDoble = estructuras.ListaDoble{Inicio: nil, Longitud: 0}
var Cola = estructuras.Cola{Primero: nil, Longitud: 0}
var Pila = estructuras.Pila{Primero: nil, Longitud: 0}

var (
	id       string
	nombre   string
	cargo    string
	password string
	opcion3  int
	nimagen  string
	capas    int
	x        int
	y        int
	color    string
	idc      string
	ide      string
	Imagen   string
	longitud int
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
		verificacion()
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
			var ruta2 string
			fmt.Println("Ingrese la ruta del archivo")
			fmt.Println("")
			fmt.Scanln(&ruta2)
			LeerArchivoImagen(ruta2)
		case 3:
			var ruta3 string
			fmt.Println("Ingrese la ruta del archivo")
			fmt.Println("")
			fmt.Scanln(&ruta3)
			LeerArchivoCliente(ruta3)
		case 4:
			var ruta4 string
			fmt.Println("Ingrese la ruta del archivo")
			fmt.Println("")
			fmt.Scanln(&ruta4)
			LeerArchivoCola(ruta4)
		case 5:
			ListaDoble.Reporte()
			ListaCircular.ReporteClientes()
			ListaSimple.ReporteEmpleados()
			Cola.Graficar()
			Pila.Graficar()
		case 6:
			fmt.Println("Cerrando Aplicacion....")
			main()

		default:
			fmt.Println("Opción inválida")
		}

	}

}

func menu_empleado() {
	opcion := 0
	for opcion != 3 {
		fmt.Println("-------------EDD Creative " + id + "-----------")
		fmt.Println("1. Ver Imagenes Cargadas")
		fmt.Println("2. Realizar Pedido")
		fmt.Println("3. Cerrar Sesión")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Println("Escriba el nombre de la imagen: ")
		case 2:
			Cola.MostrarPrimero()
			RealizarPedido(id)
		default:

		}
	}
}

func RealizarPedido(id string) {
	for {
		idc = Cola.ObtenerId()
		nombre = Cola.ObtenerNombre()
		longitud = Cola.ObtenerLongitud()
		if longitud != 0 {
			fmt.Println("Atendiendo al cliente con id: " + idc)
			fmt.Println("Atendiendo al cliente: " + nombre)
			if strings.ToUpper(idc) == "X" {
				for {
					var nuevoid int
					_ = nuevoid
					nuevoid = rand.Intn(9999)
					validacion := ListaCircular.Validar(strconv.Itoa(nuevoid))
					if validacion == true {
						// " "
					} else {
						ListaDoble.Mostrar()
						var opcion int
						fmt.Println("Elija una opción: ")
						fmt.Scanln(&opcion)
						nimagen = ListaDoble.Validarimagen(strconv.Itoa(opcion))
						ListaCircular.Insertar(idc, nombre)
						Pila.Push(idc, nimagen, id)
						Cola.Descolar()
						fmt.Println("Cliente agregado: ", nombre, "con id ", nuevoid)
						break
					}
				}
			} else {
				var opcion int
				validacion := ListaCircular.Validar(idc)
				if validacion == true {
					ListaDoble.Mostrar()
					fmt.Println("Elija una opción: ")
					fmt.Scanln(&opcion)
					nimagen = ListaDoble.Validarimagen(strconv.Itoa(opcion))
					Pila.Push(idc, nimagen, id)
					Cola.Descolar()
				} else {
					ListaDoble.Mostrar()
					fmt.Println("Elija una opción: ")
					fmt.Scanln(&opcion)
					nimagen = ListaDoble.Validarimagen(strconv.Itoa(opcion))
					ListaCircular.Insertar(idc, nombre)
					Pila.Push(idc, nimagen, id)
					Cola.Descolar()
				}

			}
		} else {
			fmt.Println("No hay clientes en cola")
			break
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
		ListaSimple.Insertar(linea[0], linea[1], linea[2], linea[3])
		fmt.Println("nombre: ", linea[1]+" Cargo: "+linea[2], " id: ", linea[0])
	}
}

func LeerArchivoImagen(ruta2 string) {
	file, err := os.Open(ruta2)
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

		str, err := strconv.Atoi(linea[1])
		ListaDoble.Insertar(linea[0], str)
		fmt.Println("Imagen: ", linea[0]+" Capas: "+linea[1])
	}
}

func LeerArchivoCliente(ruta3 string) {
	file, err := os.Open(ruta3)
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
		ListaCircular.Insertar(linea[0], linea[1])
		fmt.Println("Id: ", linea[0]+" Nombre: "+linea[1])
	}
}

func LeerArchivoCola(ruta4 string) {
	file, err := os.Open(ruta4)
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
		Cola.Encolar(linea[0], linea[1])
		fmt.Println("Id: ", linea[0]+" Nombre: "+linea[1])
	}
}

/*func ExisteId(id string) bool {
	if ListaCircular.RecorrerClientes(id, nombre) != nil {
		Cola.RecorrerCola(id, nombre)
		if Cola.RecorrerCola().Id == ListaCircular.RecorrerClientes().Id {
			ListaCircular.Insertar(Cola.Enlistar().Id, Cola.Enlistar().Nombre)
			return true
		}
	} else {
		return false
	}
}*/

func InsertarImagen() {
	//ListaDoble.Insertar(imagen, capas)
}

func verificacion() {

	if ListaSimple.Recorrer(id, password) != nil {
		fmt.Println("Se ha iniciado sesión")
		fmt.Println("Bienvenido" + id)
	} else {
		fmt.Println(
			"Credenciales incorrectas",
		)
	}

	menu_empleado()
}

func matrizz() {
	matriz := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	matriz_csv := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	matriz.Insertar_Elemento(x, y, color)
	//matriz.Reporte()
	imagen := ListaDoble.Recorrer(nimagen, capas).Imagen
	//archivo := "body.csv"
	//matriz_csv.LeerArchivo("csv/" + imagen + "/" + archivo)
	//matriz_csv.Reporte()
	matriz_csv.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
	matriz_csv.GenerarImagen(imagen)
}
