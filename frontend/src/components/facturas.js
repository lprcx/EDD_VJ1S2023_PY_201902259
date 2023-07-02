import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Factura = () => {
    const [imagen, setImagen] = useState('https://static.vecteezy.com/system/resources/thumbnails/014/612/267/small/hand-drawn-dotted-line-for-border-decoration-underline-a-cute-quote-png.png')
    const [imagen2, setImagen2] = useState('https://th.bing.com/th/id/R.49b348aeab77c0e11ec15754f00f64d0?rik=la5X6Djf6FBVqw&pid=ImgRaw&r=0')
    const idEmpleado = localStorage.getItem("empleado")
    const [facturas, setFacturas] = useState([])
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleado","_self")
    }

    useEffect(() => {
        peticion()
    },[])

    const peticion = () => {
        fetch('http://localhost:3001/facturaempleado',{
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        console.log(data.factura)
        setFacturas(data.factura) 
    }

    return(
        <div className="form-signin2">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal" style={{ color: 'green' }}>Facturas Generadas <br/> ID Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <table className="table table-bordered">
                        <thead>
                            <tr>
                                <th scope="col">#</th>
                                <th scope="col">ID Cliente</th>
                                <th scope="col">ID Factura</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                facturas.map((element, j) => {
                                    if (element.Id_Cliente != '') {
                                        return <>
                                        <tr key={"fact"+j}>
                                            <th scope="row">{j+1}</th>
                                            <td>{element.Id_Cliente}</td>
                                            <td>{element.Id_Factura}</td>
                                        </tr>
                                    </>
                                    }
                                })
                            }
                        </tbody>
                    </table>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <center><img src={imagen} width="300" height="100" alt='some value' /></center>
                  </form>
            </div>
            <center><img src={imagen2} width="200" height="200" alt='some value' /></center>
          </div>
    );
}