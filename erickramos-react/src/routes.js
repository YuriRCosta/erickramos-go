import React from "react";
import { useRoutes } from "react-router-dom";
import Login from "./pages/Login";
import RetentorComando from "./pages/RetentorComando";
import Valvula from "./pages/Valvula";
import Selo from "./pages/Selo";
import Junta from "./pages/Junta";
import Comando from "./pages/Comando";
import Cabecote from "./pages/Cabecote";

const Routess = () => {
	return useRoutes([
		{ path: "/", element: <Login /> },
		{ path: "/retentoresComando", element: <RetentorComando /> },
		{ path: "/retentoresValvula", element: <RetentorValvula /> },
		{ path: "/valvulas", element: <Valvula /> },
		{ path: "/selos", element: <Selo /> },
		{ path: "/juntas", element: <Junta /> },
		{ path: "/comandos", element: <Comando /> },
		{ path: "/cabecotes", element: <Cabecote /> },
	]);
};

export default Routess;
