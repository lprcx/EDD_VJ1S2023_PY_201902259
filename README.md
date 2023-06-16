# EDD_VJ1S2023_PY_201902259
# EDD Creative
## FASE 1

### Objetivo General
Aplicar los conocimientos del curso Estructuras de Datos en el desarrollo de
las diferentes estructuras de datos y los diferentes algoritmos de
manipulación de la información en ellas.
#### Objetivos Específicos
1. Utilizar el lenguaje Go para implementar estructuras de datos lineales
2. Utilizar la herramienta Graphviz para graficar las estructuras de datos.
3. Definir e implementar algoritmos de ordenamiento, búsqueda e inserción para
las listas enlazadas.


### Manual Técnico

#### Funciones más relevantes

##### Estructuras/listasimple.go
###### Función Insertar

```go
func (l *ListaEmpleados) Insertar(id string, nombre string, cargo string, password string) {
	nuevoEmpleado := &Empleado{Id: id, Nombre: nombre, Cargo: cargo, Password: password}
	if l.estaVacia() {
		l.Inicio = &NodoLista{empleado: nuevoEmpleado, siguiente: nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoLista{empleado: nuevoEmpleado, siguiente: nil}
		l.Longitud++
	}
}
```
En listasimple.go se pueden encontrar las funciones para añadir a un empleado a una lista simple de empleados, así como la estructura de una lista simple y el reporte graficado de la misma.

###### Función Recorrer
```go
func (l *ListaEmpleados) Recorrer(id string, password string) *Empleado {
	if l.estaVacia() {
		return nil
	} else {
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			if id == aux.empleado.Id && password == aux.empleado.Password {
				return aux.empleado
			}
			aux = aux.siguiente
		}
		return nil
	}
}
```
La función Recorrer() se realizó para leer el id y la contraseña de la lista de empleados para luego comparar esos valores y permitir la validación de usuarios.


##### Estructuras/listadoble.go
###### Función Reporte

```go
func (l *ListaDoble) Reporte() {
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.nimagen.Imagen + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
```
La función reporte realiza un gráfico de la lista doble, haciendo que el nodo n apunte a un nodo n+1, seguido del nodo n+1 apuntando al nodo n, para generar de esa manera un grafico de lista doble enlazada.