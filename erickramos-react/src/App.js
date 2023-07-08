// src/App.js
import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import RetentorComando from "./pages/RetentorComando";
import RetentorValvula from "./pages/RetentorValvula";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/retentoresComando" element={<RetentorComando />} />
        <Route path="/retentoresValvula" element={<RetentorValvula />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
