import * as React from "react"; 
import { Oval } from 'react-loader-spinner'
import * as ReactDOM from "react-dom/client";
import './App.css';
import {connect, getClientName} from "./socket.js";
import {useEffect, useState} from 'react';

import * as SockJS from 'sockjs-client';
import * as Stomp from 'stompjs';

export default function Node() {


    useEffect(() => {connect(Stomp)}, [])
    
    return (
        <>
            <div className="socket-header">Viewing Logs For: {getClientName().split("/").pop()}</div>
            <div id="messages" style={{display: "flex", flexFlow: "column"}}> </div>
            <div style={{ margin: "20px"}}>
             <Oval
                  visible={true}
                  height="30"
                  width="30"
                  color="black"
                  ariaLabel="oval-loading"
                  wrapperStyle={{}}
                  wrapperClass=""
              /> 
            </div>
        </>
    )
}
