package PrimeraFase

import (
	estructuras "EDD_VJ1S2023_PY_201902259/PrimeraFase/Estructuras"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gorilla/mux"
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
	for opcion != 4 {
		fmt.Println("-------------EDD Creative " + id + "-----------")
		fmt.Println("1. Ver Imagenes Cargadas")
		fmt.Println("2. Realizar Pedido")
		fmt.Println("3. Reporte de Capas")
		fmt.Println("4. Cerrar Sesión")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			matrizz()
		case 2:
			Cola.MostrarPrimero()
			RealizarPedido(id)
		case 3:
			ReporteCapa()
		default:

		}
	}
}

func ReporteCapa() {
	opcion := 0
	ListaDoble.Mostrar()
	fmt.Println("Elija una opción: ")
	fmt.Scanln(&opcion)
	nimagen = ListaDoble.Validarimagen(strconv.Itoa(opcion))
	var listacapas = estructuras.NewListaSimpleC()
	var matriz = &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	matriz.LeerInicialCapas("csv/"+nimagen+"/inicial.csv", nimagen, listacapas)
	opcion1 := 0
	listacapas.ListarCapa()
	fmt.Println("Elija una opción: ")
	fmt.Scanln(&opcion1)
	nombrecapa := listacapas.BuscarCapa(strconv.Itoa(opcion1))
	matriz = &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	matriz.LeerCapa2("csv/"+nimagen+"/inicial.csv", nimagen, nombrecapa)
	matriz = &estructuras.Matriz{Raiz: nil}
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
						ListaCircular.Insertar(strconv.Itoa(nuevoid), nombre)
						Pila.Push(strconv.Itoa(nuevoid), nimagen, id)
						Cola.Descolar()
						fmt.Println("Cliente agregado: ", nombre, "con id ", nuevoid)
						contenido := estructuras.ArchivoJSON(&Pila)
						estructuras.CrearArchivo()
						estructuras.EscribirArchivo(contenido)
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
					contenido := estructuras.ArchivoJSON(&Pila)
					estructuras.CrearArchivo()
					estructuras.EscribirArchivo(contenido)
				} else {
					ListaDoble.Mostrar()
					fmt.Println("Elija una opción: ")
					fmt.Scanln(&opcion)
					nimagen = ListaDoble.Validarimagen(strconv.Itoa(opcion))
					ListaCircular.Insertar(idc, nombre)
					Pila.Push(idc, nimagen, id)
					Cola.Descolar()
					contenido := estructuras.ArchivoJSON(&Pila)
					estructuras.CrearArchivo()
					estructuras.EscribirArchivo(contenido)
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
		//fmt.Println("nombre: ", linea[1]+" Cargo: "+linea[2], " id: ", linea[0])
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
		//fmt.Println("Imagen: ", linea[0]+" Capas: "+linea[1])
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
		//fmt.Println("Id: ", linea[0]+" Nombre: "+linea[1])
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
		//fmt.Println("Id: ", linea[0]+" Nombre: "+linea[1])
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
	opcion := 0
	var matriz = estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	//matriz.Insertar_Elemento(x, y, color)
	//matriz.Reporte()
	//imagen := ListaDoble.Recorrer(nimagen, capas).Imagen
	//archivo := "body.csv"
	//matriz_csv.LeerArchivo("csv/" + imagen + "/" + archivo)
	//matriz_csv.Reporte()
	ListaDoble.Mostrar()
	fmt.Println("Elija una opción: ")
	fmt.Scanln(&opcion)
	nimagen = ListaDoble.Validarimagen(strconv.Itoa(opcion))
	matriz.LeerInicial("csv/"+nimagen+"/inicial.csv", nimagen)
	matriz.GenerarImagen(nimagen)
	matriz = estructuras.Matriz{Raiz: nil}
	fmt.Println("Imagen generada con éxito " + nimagen)
}

//Segunda fase

type RespuestaImagen struct {
	Imagenbase64 string
	Nombre       string
}

var arbol *estructuras.Arbol

func mainArbol() {
	arbol = &estructuras.Arbol{Raiz: nil}
	r := mux.NewRouter()
	/*Devolver algo*/
	r.HandleFunc("/", MostrarArbol).Methods("GET")
	/*Recibe un valor del frontend*/
	r.HandleFunc("/agregar-arbol", AgregarArbol).Methods("POST")
	r.HandleFunc("/reporte-arbol", MandarReporte).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", r))
}

func MostrarArbol(w http.ResponseWriter, req *http.Request) {
	/*Esto nos verifica que le estamos enviando al servidor una respuesta de tipo JSON*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&arbol)
}

func AgregarArbol(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	var nuevoNodo estructuras.NodoArbol
	if err != nil {
		fmt.Fprintf(w, "No valido")
	}
	json.Unmarshal(reqBody, &nuevoNodo)
	arbol.InsertarElemento(nuevoNodo.Valor)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoNodo)
}

func MandarReporte(w http.ResponseWriter, req *http.Request) {
	arbol.Graficar()
	var imagen RespuestaImagen = RespuestaImagen{Nombre: "arbolAVL.jpg"}
	/*INICIO*/
	imageBytes, err := ioutil.ReadFile(imagen.Nombre)
	if err != nil {
		fmt.Fprintf(w, "Imagen No Valida")
		return
	}
	// Codifica los bytes de la imagen en base64
	imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)

	/*data:image/jpg;base64,ABC*/
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imagen)
}

// clase 10 y 11
// clase 11
type Persona struct {
	Identificador string
	Password      string
}

/*
{
	Identificador: "usuario1"
	Password: "1234"
}
*/

var Tablahash *estructuras.TablaHash

func mainHash() {
	Tablahash = &estructuras.TablaHash{Capacidad: 5, Utilizacion: 0}
	Tablahash.NewTablaHash()
	app := fiber.New()
	app.Use(cors.New())
	// app.handlerFunc("/ruta",nombreFuncion).metodo("GET")
	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(&fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/comprobar-usuario", func(c *fiber.Ctx) error {
		var usuario Persona
		c.BodyParser(&usuario)
		fmt.Println(usuario)
		if usuario.Identificador == "2017" && usuario.Password == "2017" {
			return c.JSON(&fiber.Map{
				"status": "OK",
			})
		}
		return c.JSON(&fiber.Map{
			"estado": "NO",
		})
	})

	app.Get("/obtener-tabla", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Datos": Tablahash,
		})
	})

	/*8793, 4593, 7893, 2356*/
	app.Post("/agregar-tabla", func(c *fiber.Ctx) error {
		var nuevo estructuras.NodoHash
		c.BodyParser(&nuevo)
		Tablahash.Insertar(nuevo.Id_Cliente, nuevo.Id_Factura)
		return c.JSON(&fiber.Map{
			"status": "OK",
		})
	})

	app.Listen(":3001")
}

//clase 15

var ListaBloque *estructuras.BlockChain

func mainnbloque() {
	ListaBloque = &estructuras.BlockChain{BloqueCreados: 0}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/agregar-bloque", func(c *fiber.Ctx) error {
		var nuevoNodo estructuras.NodoBloquePeticion
		c.BodyParser(&nuevoNodo)
		ListaBloque.InsertarBloque(nuevoNodo.Timestamp, nuevoNodo.Biller, nuevoNodo.Customer, nuevoNodo.Payment)
		return c.JSON(&fiber.Map{
			"Data": nuevoNodo,
		})
	})

	app.Get("/tablafacturas", func(c *fiber.Ctx) error {
		factura := ListaBloque.ArregloFacturas()
		return c.JSON(&fiber.Map{
			"message": "OK",
			"data":    factura,
		})
	})
	app.Listen(":3001")
}
