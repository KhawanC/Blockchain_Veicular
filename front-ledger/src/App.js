import React from 'react';
import { Root } from "./Routes/Root"
import "./global-style.css"
import { Navvbar } from './Components/Navbar';

export function App() {
  return (
    <div className="App">
      <Navvbar/>
      <Root/>
    </div>
  );
}
