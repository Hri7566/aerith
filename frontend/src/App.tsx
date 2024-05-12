import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet, LoadMIDI, RenderMIDI, UnloadMIDI } from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [statusText, setStatusText] = useState("<status>");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    function render() {
        RenderMIDI().then(status => {
            setStatusText(status);
        });
    }

    function load() {
        LoadMIDI().then(status => {
            setStatusText(status);
        });
    }

    function unload() {
        UnloadMIDI().then(status => {
            setStatusText(status);
        });
    }

    return (
        <div id="App">
            <p id="status">{statusText}</p>
            <button onClick={load}>Load</button>
            <button onClick={unload}>Unload</button>
            <button onClick={render}>Render</button>
        </div>
    )
}

export default App
