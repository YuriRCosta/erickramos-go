import React from "react";
import {useRoutes} from "react-router-dom";
import Login from "./pages/Login";
import RetentorComando from "./pages/RetentorComando";

const Routess = () => {
    return useRoutes([
        {path: '/', element: <Login/>},
        {path: '/retentoresComando', element: <RetentorComando/>}
        ]
    );
}

export default Routess;