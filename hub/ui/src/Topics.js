import logo from './logo.svg';
import { Oval } from 'react-loader-spinner'
import './App.css';
import {listTopics} from "./socket.js";
import {useEffect, useState} from 'react';

function Topics() {
    
    const [topics, setTopics] = useState([]);

    useEffect(
        () => 
            { 
                async function fetchTopics() {
                    listTopics(function(data) {
                        console.log(data)
                        setTopics(data)
                    });
                }
                fetchTopics();
            }, 
    [])

    return (
        <div className="App">
            <div className="container">
            <div style={{fontWeight: "bold", marginBottom: "20px"}}>Servers</div>
            { 
                topics.length == 0 ? 
                 <Oval
                      visible={true}
                      height="30"
                      width="30"
                      color="black"
                      ariaLabel="oval-loading"
                      wrapperStyle={{}}
                      wrapperClass=""
                  /> : 
                Object.keys(topics).map(key =>
                    <div key={topics[key]} style={{marginBottom: "10px"}}>
                    <a href={"/node?node=" + topics[key]} >{topics[key]}</a>
                    </div>
                ) 
            }
            </div>
            <img src={"./sms.png"} style={{width: "125px", height: "50px", position: "fixed", top: "40px", left: "20px"}}/>
        </div>
    );
}

export default Topics;
