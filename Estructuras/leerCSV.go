package estructuras

func ArchivoJSON(p *Pila) string {
	contenido := "{\n"
	contenido += "\t\"pedidos\": [\n"
	aux := p.Primero
	for aux.Siguiente != nil {
		contenido += "\t\t{\n"
		if aux.idcliente == "X" {
			contenido += "\t\t\t\"id_cliente\": \"" + (aux.idcliente) + "" + "\", \n"
			contenido += "\t\t\t\"imagen\": \"" + (aux.Imagen) + " " + "\" \n"
			contenido += "\t\t},\n"
		} else {
			contenido += "\t\t\t\"id_cliente\": " + (aux.idcliente) + ", \n"
			contenido += "\t\t\t\"imagen\": \"" + (aux.Imagen) + " " + "\" \n"
			contenido += "\t\t},\n"
		}
		aux = aux.Siguiente
	}
	//esto es para el ultimo elemento
	contenido += "\t\t{\n"
	contenido += "\t\t\t\"id_cliente\": " + (aux.idcliente) + ", \n"
	contenido += "\t\t\t\"imagen\": \"" + (aux.Imagen) + " " + "\" \n"
	contenido += "\t\t}\n"
	contenido += "\t]\n"
	contenido += "}"
	return contenido
}

func Generarjson(p *Pila) {
	CrearArchivo()
	EscribirArchivo(ArchivoJSON(p))
}
