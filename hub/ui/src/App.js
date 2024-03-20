import logo from './logo.svg';
import { Oval } from 'react-loader-spinner'
import './App.css';
import {listClients} from "./socket.js";
import {useEffect, useState} from 'react';

function App() {
    
    const [clients, setClients] = useState([]);

    useEffect(
        () => 
            { 
                async function fetchClients() {
                    listClients(function(data) {
                        console.log(data);
                        setClients(data)
                    });
                }
                fetchClients();
            }, 
    [])

    return (
        <div className="App">
            <div className="container">
            <div style={{fontWeight: "bold", marginBottom: "20px"}}>Connected MAS Cluster Clients</div>
            { 
                clients.length == 0 ? 
                 <Oval
                      visible={true}
                      height="30"
                      width="30"
                      color="black"
                      ariaLabel="oval-loading"
                      wrapperStyle={{}}
                      wrapperClass=""
                  /> : 
                Object.keys(clients).map(key => 
                    <div key={clients[key]} style={{marginBottom: "10px"}}>
                    <a href={"/topics?name=" + clients[key].nodeName} >{clients[key].nodeName}</a>
                    </div>
                ) 
            }
            </div>
            <img src={"./sms.png"} style={{width: "125px", height: "50px", position: "fixed", top: "40px", left: "20px"}}/>
        </div>
    );
}

export default App;
