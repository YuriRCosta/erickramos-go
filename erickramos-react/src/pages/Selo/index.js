import React, { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import api from "../../services/api";

export default function Selo() {
    const [selos, setSelos] = useState([]);

    const [medidaPesquisa, setMedidaPesquisa] = useState("");
    const [nomePesquisa, setNomePesquisa] = useState("");
    const [modalVisible, setModalVisible] = useState(false);

    const accessToken = localStorage.getItem("token");

    const openModal = () => {
        setModalVisible(true);
    };

    const closeModal = () => {
        setModalVisible(false);
    };

    async function carregarSelos() {
        const response = await api.get("selos", {
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
        });
        setSelos(response.data);
    }

    const handleMedidaChange = (e) => {
        setMedidaPesquisa(e.target.value);
    };

    const handleNomeChange = (e) => {
        setNomePesquisa(e.target.value);
    };

    const handlePesquisar = () => {
        if (medidaPesquisa) {
            api.get(`/selos/medida/${medidaPesquisa}`, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    const resultado = response.data;
                    console.log(resultado);
                    setSelos(resultado);
                })
                .catch((error) => {
                    console.error(error);
                });
        } else if (nomePesquisa) {
            api.get(`/selos/nome/${nomePesquisa}`, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    const resultado = response.data;
                    console.log(resultado);
                    setSelos(resultado);
                })
                .catch((error) => {
                    console.error(error);
                });
        } else {
            console.log("Informe um critério de pesquisa válido.");
        }
    };

    async function deleteSelo(id) {
        const confirmDelete = window.confirm("Tem certeza que deseja deletar?");

        if (!confirmDelete) {
            return;
        }
        try {
            await api.delete(`selos/${id}`, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            });
            setSelos(selos.filter((selo) => selo.id !== id));
        } catch (e) {
            alert("Falha ao deletar selo, tente novamente.");
        }
    }

    const [nome, setNome] = useState("");
    const [preco, setPreco] = useState("");
    const [medida, setMedida] = useState("");
    const [qtd_estoque, setQtdEstoque] = useState("");

    const handleNomeChangeModal = (e) => {
        setNome(e.target.value);
    };

    const handlePrecoChangeModal = (e) => {
        setPreco(e.target.value);
    };

    const handleMedidaChangeModal = (e) => {
        setMedida(e.target.value);
    };

    const handleQtdEstoqueChangeModal = (e) => {
        setQtdEstoque(e.target.value);
    };

    const handleSave = () => {
        if (nome && preco && medida && qtd_estoque) {
            const novoSelo = {
                nome: nome,
                preco: parseFloat(preco),
                medida: medida,
                qtd_estoque: parseInt(qtd_estoque),
            };

            api.post("/selos", novoSelo, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    setSelos([...selos, response.data]);

                    setNome("");
                    setPreco("");
                    setMedida("");
                    setQtdEstoque("");

                    closeModal();
                })
                .catch((error) => {
                    console.error(error);
                });
        } else {
            console.log("Preencha todos os campos do formulário.");
        }
    };

    const handleCancel = () => {
        setNome("");
        setPreco("");
        setMedida("");
        setQtdEstoque("");
        closeModal();
    };

    useEffect(() => {
        carregarSelos();
    }, []);

    return (
        <div>
            <Navbar />
            <div className="flex items-center bg-white justify-center h-screen">
                <div className="max-w-fit w-full p-6 bg-neutral-500 rounded-lg shadow-md text-center">
                    <h1 className="text-2xl font-semibold mb-4">Selos</h1>

                    <div className="flex flex-col items-center">
                        <input
                            type="text"
                            value={medidaPesquisa}
                            onChange={handleMedidaChange}
                            placeholder="Pesquisar por medida"
                            className="px-4 mb-2 text-black py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
                        />
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
                                onClick={carregarSelos}
                                className="px-4 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none"
                            >
                                Mostrar Todos
                            </button>
                        </div>
                    </div>

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
                                    Medida
                                </th>
                                <th className="py-2 text-center px-4 border-b">
                                    Qtd Estoque
                                </th>
                                <th className="py-2 text-center px-4">
                                    Excluir
                                </th>
                                <th className="py-2 text-center px-4">
                                    Editar
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {selos == null ? (
                                <tr>
                                    <td
                                        colSpan="6"
                                        className="py-4 text-center"
                                    >
                                        Não há selos cadastrados.
                                    </td>
                                </tr>
                            ) : (
                                selos.map((selo) => (
                                    <tr key={selo.id}>
                                        <td className="hidden">{selo.id}</td>
                                        <td className="py-2 text-center px-4 border-b">
                                            {selo.nome}
                                        </td>
                                        <td className="py-2 text-center px-4 border-b">
                                            R$ {selo.preco}
                                        </td>
                                        <td className="py-2 text-center px-4 border-b">
                                            {selo.medida}
                                        </td>
                                        <td className="py-2 text-center px-4 border-b">
                                            {selo.qtd_estoque}
                                        </td>
                                        <button
                                            className="p-2 text-red-500 hover:text-red-700 focus:outline-none"
                                            onClick={() => deleteSelo(selo.id)}
                                        >
                                            <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                className="w-5 h-5"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke="currentColor"
                                                strokeWidth="2"
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                            >
                                                <path d="M3 6l3 15h12l3-15H3zm10 3v9M8 6h8M9 3h6M4 6h16" />
                                            </svg>
                                        </button>
                                    </tr>
                                ))
                            )}
                        </tbody>
                    </table>
                    <div className="flex justify-center mt-4">
                        <button
                            onClick={openModal}
                            className="px-4 py-2 text-white bg-green-500 rounded-lg hover:bg-green-600 focus:outline-none"
                        >
                            Adicionar Selo
                        </button>
                    </div>
                </div>
                {modalVisible && (
                    <div className="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-75">
                        <div className="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-75">
                            <div className="bg-white w-1/2 p-6 rounded-lg shadow-lg">
                                <h2 className="text-xl font-semibold mb-4">
                                    Novo Selo
                                </h2>

                                <div>
                                    <label className="block mb-2">
                                        Nome:
                                        <input
                                            type="text"
                                            value={nome}
                                            onChange={handleNomeChangeModal}
                                            className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 w-full"
                                        />
                                    </label>
                                    <label className="block mb-2">
                                        Preço:
                                        <input
                                            type="number"
                                            value={preco}
                                            onChange={handlePrecoChangeModal}
                                            className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 w-full"
                                        />
                                    </label>
                                    <label className="block mb-2">
                                        Medida:
                                        <input
                                            type="text"
                                            value={medida}
                                            onChange={handleMedidaChangeModal}
                                            className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 w-full"
                                        />
                                    </label>
                                    <label className="block mb-2">
                                        Quantidade em Estoque:
                                        <input
                                            type="number"
                                            value={qtd_estoque}
                                            onChange={
                                                handleQtdEstoqueChangeModal
                                            }
                                            className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 w-full"
                                        />
                                    </label>
                                </div>

                                <div className="flex justify-end mt-4">
                                    <button
                                        onClick={handleCancel}
                                        className="px-4 py-2 text-white bg-red-500 rounded-lg hover:bg-red-600 focus:outline-none mr-2"
                                    >
                                        Cancelar
                                    </button>
                                    <button
                                        onClick={handleSave}
                                        className="px-4 py-2 text-white bg-green-500 rounded-lg hover:bg-green-600 focus:outline-none"
                                    >
                                        Salvar
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}
