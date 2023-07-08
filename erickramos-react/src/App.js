// src/App.js
import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import RetentorComando from "./pages/RetentorComando";
import RetentorValvula from "./pages/RetentorValvula";
import Valvula from "./pages/Valvula";
import Selo from "./pages/Selo";
import Junta from "./pages/Junta";
import Comando from "./pages/Comando";
import Cabecote from "./pages/Cabecote";

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Login />} />
				<Route
					path="/retentoresComando"
					element={<RetentorComando />}
				/>
				<Route
					path="/retentoresValvula"
					element={<RetentorValvula />}
				/>
				<Route path="valvulas" element={<Valvula />} />
				<Route path="selos" element={<Selo />} />
				<Route path="juntas" element={<Junta />} />
				<Route path="comandos" element={<Comando />} />
				<Route path="cabecotes" element={<Cabecote />} />
			</Routes>
		</BrowserRouter>
	);
}

export default App;
