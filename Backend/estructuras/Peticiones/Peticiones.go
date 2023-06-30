package peticiones

type Login struct {
	Username string
	Password string
}

type Archivo struct {
	Nombre string
}

type Pedido struct {
	Id_Cliente    int    `json:"id_cliente"`
	Nombre_Imagen string `json:"nombre_imagen"`
}

type ArbolPeticion struct {
	Pedidos []Pedido
}

type RespuestaImagen struct {
	Imagenbase64 string
	Nombre       string
}

type PeticionFiltro struct {
	Tipo int
}
