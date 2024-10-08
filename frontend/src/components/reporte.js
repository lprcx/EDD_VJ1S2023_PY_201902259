import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Reportes = () => {
    const [imagen, setImagen] = useState('https://www.pngplay.com/wp-content/uploads/9/Kirby-Transparent-File.png')
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/admin","_self")
    }

    const validar = (data) =>{
        console.log(data)
        setImagen(data.imagen.Imagenbase64)
    }

    const reporteGrafo = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-grafo',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const reporteArbol = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const reporteBlockchain = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-bloque',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    return(
        <div className="form-signin1">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 style={{ color: 'pink' }}>Reportes Administrador</h1>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={reporteGrafo}>Grafo</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-info" onClick={reporteArbol}>Arbol AVL</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-warning" onClick={reporteBlockchain}>Facturas</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-danger" onClick={salir}>Salir</button></center>
                    <br/>
                    <center><img src={imagen} width="350" height="350" alt='some value' /></center>
                    <br/>
                    <p className="mt-5 mb-3 text-muted">EDD 201902259</p>
                    <br/>
                  </form>
            </div>
          </div>
    );
}