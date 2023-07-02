import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Filtros = () => {
    const [imagen, setImagen] = useState('https://www.geekmi.news/__export/1616108436990/sites/debate/img/2021/03/18/kirby_1_crop1616108354116.jpg_423682103.jpg')
    const [imagen2, setImagen2] = useState('https://static.vecteezy.com/system/resources/thumbnails/014/606/453/small/hand-drawn-dotted-line-for-border-decoration-underline-a-cute-quote-png.png')
    const [filtro, setFiltro] = useState(0)
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleado","_self")
    }

    const validar = (data) =>{
        console.log(data)
    }

    const aplicarFiltros = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/aplicarfiltro',{
            method: 'POST',
            body: JSON.stringify({
                Tipo: filtro,
                NombreImagen: ""
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const handleChange = (e) => {
        var j = parseInt(e.target.value);
        setFiltro(j)
    }
 

    return(
        <>
        <div className="form-signin1">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 style={{ color: 'pink' }}>Bienvenido {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <h4 className="h3 mb-3 fw-normal" style={{ color: 'blue' }}>Seleccione un Filtro</h4>
                    <br/>
                    <div className="col align-self-center">
                        <select className="custom-select" aria-label=".form-select-lg example" onChange={handleChange}>
                            <option value={0}>Elegir....</option>
                            <option value={1}>Filtro Negativo</option>
                            <option value={2}>Escala de Grises</option>
                            <option value={3}>Filtro Espejo X</option>
                            <option value={4}>Filtro Espejo Y</option>
                            <option value={5}>Ambos Espejos</option>
                        </select>
                    </div>
                    <br/>
                    <center><button className="w-50 btn btn-outline-warning" onClick={aplicarFiltros}>Generar Nueva Imagen</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-info" onClick={salir}>Salir</button></center>
                    <br/>
                    <p className="mt-5 mb-3 text-muted">EDD 201902259</p>
                    <center><img src={imagen2} width="300" height="100" alt='some value' /></center>
                  </form>
            </div>
            <br/>
                    <center><img src={imagen} width="350" height="350" alt='some value' /></center>
          </div>
          <br/>
        
        </>          
    );
}