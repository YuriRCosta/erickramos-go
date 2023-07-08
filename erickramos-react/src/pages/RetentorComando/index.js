import React, {useState, useEffect} from "react";
import {Link, useNavigate} from "react-router-dom";

import api from "../../services/api";

export default function RetentorComando({children}) {

    const [retentoresComando, setRetentoresComando] = useState([]);

    const accessToken = localStorage.getItem('token');

    async function fetchMoreRetentorComandos() {
        const response = await api.get('retentorComando', {
            headers: {
                Authorization: `Bearer ${accessToken}`
            },
        })
        setRetentoresComando(response.data);
    }

    useEffect(() => {
        fetchMoreRetentorComandos();
    }, []);

    return (
        <div className="flex items-center bg-white justify-center h-screen">
                <div className="max-w-md w-full p-6 bg-neutral-500 rounded-lg shadow-md text-center">
                    <h1 className="text-2xl font-semibold mb-4">Retentores de Comando</h1>

                    <div className="flex flex-col items-center">
                        <div className="flex flex-row">
                            <button
                                onClick={fetchMoreRetentorComandos}
                                className="px-4 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none"
                            >
                                Mostrar Todos
                            </button>
                        </div>
                    </div>
                    {/* Tabela */}
                    <table className="w-full">
                        <thead>
                        <tr>
                            <th className="py-2 text-center px-4 border-b">Cabecote</th>
                            <th className="py-2 text-center px-4 border-b">Preco</th>
                            <th className="py-2 text-center px-4 border-b">Qtd Estoque</th>
                        </tr>
                        </thead>
                        <tbody>
                        {retentoresComando.map((retentor) => (
                            <tr key={retentor.id}>
                                <td className="py-2 text-center px-4 border-b">{retentor.nome}</td>
                                <td className="py-2 text-center px-4 border-b">R$ {retentor.preco}</td>
                                <td className="py-2 text-center px-4 border-b">{retentor.qtdEstoque}</td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </div>
            </div>
        );
}