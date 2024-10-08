import React, {useState} from 'react';
import '../css/login.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import background from '../img/kirby.jpg'

export const Login = () => {
    const [imagen, setImagen] = useState('https://www.pngplay.com/wp-content/uploads/9/Kirby-Transparent-Background.png')
    const [user, setUser] = useState('')
    const [pass, setPass] = useState('')
    
    
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log(user + " " + pass)
        fetch('http://localhost:3001/login',{
            method: 'POST',
            body: JSON.stringify({
                Username: user,
                Password: pass
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) => {
        if(data.status == "400"){
            window.open("/admin","_self")
        }else if(data.status == "200"){
            localStorage.setItem("empleado", user)
            window.open("/empleado","_self")
        }else{
            console.log("ME dio ansiedad")
        }
    }


    return(
        <div style={{ backgroundImage: `url(${background})` }}>
            <div className="form-signin">
                <div className="text-center">
                    <form onSubmit={handleSubmit} className="card card-body">
                        <h1 className="h3 mb-3 fw-normal" style={{ color: 'blue' }}>Inicio de Sesion EDD Creative</h1>
                        <label htmlFor="inputEmail" className="visually-hidden">Usuario</label>
                        <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" required
                        onChange={e => setUser(e.target.value)} 
                        value={user}  
                        autoFocus/>
                        <br/>
                        <label htmlFor="inputPassword" className="visually-hidden">Password</label>
                        <input type="password" id="passI" className="form-control" placeholder="Password" aria-describedby="passwordHelpInline" //required 
                        onChange={e => setPass(e.target.value)}
                        value={pass} 
                        autoFocus/>
                        <br />
                        <button className="w-100 btn btn-success" type="submit">Iniciar Sesion</button>
                        <br/>
                        <center><img src={imagen} width="250" height="250" alt='some value' /></center>
                        <br/>
                        <p className="mt-5 mb-3 text-muted">EDD 201902259</p>
                        <br/>
                    </form>
                </div>
            </div>
        </div>
    );
}