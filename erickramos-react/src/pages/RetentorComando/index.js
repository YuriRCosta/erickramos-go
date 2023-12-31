import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";

import api from "../../services/api";
import Navbar from "../../components/Navbar";

export default function RetentorComando({ children }) {
    const [retentoresComando, setRetentoresComando] = useState([]);
    const [nomePesquisa, setNomePesquisa] = useState("");
    const accessToken = localStorage.getItem("token");

    async function fetchMoreRetentorComandos() {
        const response = await api.get("retentorComando", {
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
        });
        setRetentoresComando(response.data);
        console.log(response.data);
    }

    const handleNomeChange = (e) => {
        setNomePesquisa(e.target.value);
    };

    const handlePesquisar = () => {
        if (nomePesquisa) {
            api.get(`/retentorComando/nome/${nomePesquisa}`, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    const resultado = response.data;
                    console.log(resultado);
                    setRetentoresComando(resultado);
                })
                .catch((error) => {
                    console.error(error);
                });
        } else {
            console.log("Informe um critério de pesquisa válido.");
        }
    };

    useEffect(() => {
        fetchMoreRetentorComandos();
    }, []);

    return (
        <div>
            <Navbar />
            <div className="flex items-center bg-white justify-center h-screen">
                <div className="max-w-md w-full p-6 bg-neutral-500 rounded-lg shadow-md text-center">
                    <h1 className="text-2xl font-semibold mb-4">
                        Retentores de Comando
                    </h1>

                    <div className="flex flex-col items-center">
                        <input
                            type="text"
                            value={nomePesquisa}
                            onChange={handleNomeChange}
                            placeholder="Pesquisar por nome"
                            className="px-4 mb-2 py-2 text-black border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
                        />
                        <div className="flex flex-row">
                            <button
                                onClick={handlePesquisar}
                                className="px-4 mr-2 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none"
                            >
                                Pesquisar
                            </button>
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
                                <th className="py-2 text-center px-4 border-b">
                                    Cabecote
                                </th>
                                <th className="py-2 text-center px-4 border-b">
                                    Preco
                                </th>
                                <th className="py-2 text-center px-4 border-b">
                                    Qtd Estoque
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {retentoresComando == null ? (
                                <tr>
                                    <td
                                        colSpan="6"
                                        className="py-4 text-center"
                                    >
                                        Não há retentores de comandos
                                        cadastrados.
                                    </td>
                                </tr>
                            ) : (
                                retentoresComando.map((retentor) => (
                                    <tr key={retentor.id}>
                                        <td className="py-2 text-center px-4 border-b">
                                            {retentor.nome}
                                        </td>
                                        <td className="py-2 text-center px-4 border-b">
                                            R$ {retentor.preco}
                                        </td>
                                        <td className="py-2 text-center px-4 border-b">
                                            {retentor.qtd_estoque}
                                        </td>
                                    </tr>
                                ))
                            )}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    );
}
