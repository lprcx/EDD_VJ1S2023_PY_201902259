import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import img from '../img/kirb.jpg'

export const Empleado = () => {
    const [imagen2, setImagen2] = useState('https://static.vecteezy.com/system/resources/thumbnails/014/628/811/small/hand-drawn-dotted-line-for-border-decoration-underline-a-cute-quote-png.png')
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/","_self")
    }

    const validar = (data) =>{
        console.log(data)
    }

    const aplicarFiltros = async(e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/filtros","_self")
    }

    const generarFacturas = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol',{
        })
        .then(response => response.json())
        .then(data => validar(data));
        window.open("/factura","_self")
    }
    
    const verFacturas = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol',{
        })
        .then(response => response.json())
        .then(data => validar(data));
        window.open("/verfactura","_self")
    }

    return(
        <>
        <div className="form-signin1">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 style={{ color: 'magenta' }}>Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <center><button className="w-50 btn btn-outline-info" onClick={aplicarFiltros}>Aplicacion de Filtros</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-warning" onClick={generarFacturas}>Generar Facturas</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={verFacturas}>Ver Facturas</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-danger" onClick={salir}>Salir</button></center>
                    <center><img src={imagen2} width="300" height="100" alt='some value' /></center>
                    
                  </form>
            </div>
          </div>
        <div className="image-container">
        <img
        src={img}
        alt="kir"
        className="img-fluid"
        style={{alignSelf: 'center'}}
        />
      </div>
      </>
    );
}