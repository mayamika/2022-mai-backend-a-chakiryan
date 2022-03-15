import logo from './logo.svg';
import './App.css';
import api from "./api";
import { useState } from 'react';

function App() {
  const [text, setText] = useState("");
  api.get("/hello").then((res) => {
    setText(JSON.stringify(res.data));
  });

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          <code>{text}</code>
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
