import React, { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import api from "../../services/api";
import ReactPaginate from "react-paginate";

export default function RetentorValvula() {
    const [retentoresValvula, setRetentoresValvula] = useState([]);

    const [nomePesquisa, setNomePesquisa] = useState("");
    const [modalVisible, setModalVisible] = useState(false);
    const [todosRetentoresValvula, setTodosRetentoresValvula] = useState([]);
    const totalPages = Math.ceil(todosRetentoresValvula.length / 5);

    const accessToken = localStorage.getItem("token");

    const handlePageChange = ({ selected }) => {
        carregarRetentoresValvula(selected + 1);
    };

    const openModal = () => {
        setModalVisible(true);
    };

    const closeModal = () => {
        setModalVisible(false);
    };

    async function carregarTodosRetentoresValvula() {
        setNomePesquisa("");

        const response = await api.get(`retentorValvula`, {
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
        });
        if (response.data) {
            setTodosRetentoresValvula(response.data);
        }
    }

    async function carregarRetentoresValvula(pagina) {
        setNomePesquisa("");

        const response = await api.get(`retentorValvula/pagina/${pagina}`, {
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
        });
        setRetentoresValvula(response.data);
        carregarTodosRetentoresValvula();
    }

    const handleNomeChange = (e) => {
        setNomePesquisa(e.target.value);
    };

    const handlePesquisar = () => {
        if (nomePesquisa) {
            api.get(`/retentorValvula/nome/${nomePesquisa}`, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    const resultado = response.data;
                    console.log(resultado);
                    setRetentoresValvula(resultado);
                    setTodosRetentoresValvula(0);
                })
                .catch((error) => {
                    console.error(error);
                });
        } else {
            console.log("Informe um critério de pesquisa válido.");
        }
    };

    async function deleteRetentorValvula(id) {
        const confirmDelete = window.confirm("Tem certeza que deseja deletar?");

        if (!confirmDelete) {
            return;
        }
        try {
            await api.delete(`retentorValvula/${id}`, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            });
            carregarRetentoresValvula(1);
        } catch (e) {
            alert("Falha ao deletar retentor de valvula, tente novamente.");
        }
    }

    const editRetentorValvula = (retentorValvula) => {
        // Defina os valores iniciais do formulário de edição com base no retentor de valvula selecionado
        setId(retentorValvula.id);
        setNome(retentorValvula.nome);
        setPreco(retentorValvula.preco);
        setCodigo(retentorValvula.codigo);
        setQtdEstoque(retentorValvula.qtd_estoque);

        // Abra o modal de edição
        setModalVisible(true);
    };

    const [id, setId] = useState("");
    const [nome, setNome] = useState("");
    const [preco, setPreco] = useState("");
    const [codigo, setCodigo] = useState("");
    const [qtd_estoque, setQtdEstoque] = useState("");

    const handleNomeChangeModal = (e) => {
        setNome(e.target.value);
    };

    const handlePrecoChangeModal = (e) => {
        setPreco(e.target.value);
    };

    const handleCodigoChangeModal = (e) => {
        setCodigo(e.target.value);
    };

    const handleQtdEstoqueChangeModal = (e) => {
        setQtdEstoque(e.target.value);
    };

    const handleUpdate = () => {
        if (nome && preco && qtd_estoque) {
            const novoRetentorValvula = {
                nome: nome,
                preco: parseFloat(preco),
                codigo: codigo,
                qtd_estoque: parseInt(qtd_estoque),
            };

            api.put("/retentorValvula/" + id, novoRetentorValvula, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    carregarRetentoresValvula(1);

                    setId("");
                    setNome("");
                    setPreco("");
                    setCodigo("");
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

    const handleAdicionarEstoque = (retentorValvula) => {
        const qtd = prompt("Informe a quantidade a ser adicionada ao estoque:");

        if (qtd) {
            const novoRetentorValvula = {
                qtd_estoque: parseInt(qtd),
            };

            if (qtd > 0) {
                api.put(
                    `retentorValvula/adicionar-estoque/${retentorValvula.id}`,
                    novoRetentorValvula,
                    {
                        headers: {
                            Authorization: `Bearer ${accessToken}`,
                        },
                    }
                )
                    .then((response) => {
                        carregarRetentoresValvula(1);
                    })
                    .catch((error) => {
                        console.error(error);
                    });
            } else {
                console.log("A quantidade deve ser maior que zero.");
            }
        } else {
            console.log("A quantidade deve ser maior que zero.");
        }
    };

    const handleSave = () => {
        if (nome && preco && qtd_estoque) {
            const novoRetentorValvula = {
                nome: nome,
                preco: parseFloat(preco),
                codigo: codigo,
                qtd_estoque: parseInt(qtd_estoque),
            };

            api.post("/retentorValvula", novoRetentorValvula, {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            })
                .then((response) => {
                    carregarTodosRetentoresValvula();
                    carregarRetentoresValvula(1);

                    setId("");
                    setNome("");
                    setPreco("");
                    setCodigo("");
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
        setId("");
        setNome("");
        setPreco("");
        setCodigo("");
        setQtdEstoque("");
        closeModal();
    };

    const formatPrice = (price) => {
        const formatter = new Intl.NumberFormat("pt-BR", {
            style: "currency",
            currency: "BRL",
        });

        return formatter.format(price);
    };

    useEffect(() => {
        carregarTodosRetentoresValvula();
        carregarRetentoresValvula(1);
    }, []);

    return (
        <div>
            <Navbar />
            <div className="flex items-center bg-gray-100 min-h-screen">
                <div className="max-w-4xl mx-auto p-6 bg-white rounded-lg shadow-md">
                    <h1 className="text-3xl font-bold text-center mb-6">
                        Gerenciar Retentores de Valvula
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
                                onClick={() => carregarRetentoresValvula(1)}
                                className="px-4 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none"
                            >
                                Mostrar Todos
                            </button>
                        </div>
                    </div>

                    <table className="w-full mt-4 bg-white shadow-md rounded-lg overflow-hidden">
                        <thead className="bg-gray-200">
                            <tr>
                                <th className="py-3 px-4 text-center border-b">
                                    Cabecote
                                </th>
                                <th className="py-3 px-4 text-center border-b">
                                    Preco
                                </th>
                                <th className="py-3 px-4 text-center border-b">
                                    Codigo
                                </th>
                                <th className="py-3 px-4 text-center border-b">
                                    Qtd Estoque
                                </th>
                                <th className="py-3 px-4 text-center border-b">
                                    Excluir
                                </th>
                                <th className="py-3 px-4 text-center border-b">
                                    Editar
                                </th>
                                <th className="py-3 px-4 text-center border-b">
                                    Add Estoque
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {retentoresValvula == null ? (
                                <tr>
                                    <td
                                        colSpan="6"
                                        className="py-4 text-center"
                                    >
                                        Não há retentores de valvula
                                        cadastrados.
                                    </td>
                                </tr>
                            ) : (
                                retentoresValvula.map((retentorValvula) => (
                                    <tr key={retentorValvula.id}>
                                        <td className="hidden">
                                            {retentorValvula.id}
                                        </td>
                                        <td className="py-3 px-4 text-center border-b">
                                            {retentorValvula.nome}
                                        </td>
                                        <td className="py-3 px-4 text-center border-b">
                                            {formatPrice(retentorValvula.preco)}
                                        </td>
                                        <td className="py-3 px-4 text-center border-b">
                                            {retentorValvula.codigo}
                                        </td>
                                        {retentorValvula.qtd_estoque <= 5 ? (
                                            <td className="py-3 px-4 text-center border-b bg-red-400">
                                                {retentorValvula.qtd_estoque}
                                            </td>
                                        ) : (
                                            <td className="py-3 px-4 text-center border-b">
                                                {retentorValvula.qtd_estoque}
                                            </td>
                                        )}
                                        <td className="py-3 px-4 text-center border-b">
                                            <button
                                                className="p-2 text-red-500 hover:text-red-700 focus:outline-none"
                                                onClick={() =>
                                                    deleteRetentorValvula(
                                                        retentorValvula.id
                                                    )
                                                }
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
                                        </td>
                                        <td className="py-3 px-4 text-center border-b">
                                            <button
                                                className="p-2 text-blue-500 hover:text-blue-700 focus:outline-none"
                                                onClick={() =>
                                                    editRetentorValvula(
                                                        retentorValvula
                                                    )
                                                }
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
                                                    <path d="M12 20h9"></path>
                                                    <path d="M5 20h5"></path>
                                                    <path d="M4 10h3"></path>
                                                    <path d="M7.5 3L12 8.5"></path>
                                                    <path d="M12 8.5L16.5 3"></path>
                                                    <path d="M16.5 21L7.5 21"></path>
                                                    <path d="M7.5 3C9.15685 4.65685 9.15685 7.34315 7.5 9"></path>
                                                </svg>
                                            </button>
                                        </td>
                                        <td className="py-3 px-4 text-center text-xl border-b">
                                            <button
                                                onClick={() =>
                                                    handleAdicionarEstoque(
                                                        retentorValvula
                                                    )
                                                }
                                                className="px-4 py-2 text-white bg-green-500 rounded-lg hover:bg-green-600 focus:outline-none"
                                            >
                                                +
                                            </button>
                                        </td>
                                    </tr>
                                ))
                            )}
                        </tbody>
                    </table>
                    <ReactPaginate
                        previousLabel={"Anterior"}
                        nextLabel={"Próximo"}
                        breakLabel={"..."}
                        pageCount={totalPages}
                        marginPagesDisplayed={2}
                        pageRangeDisplayed={5}
                        onPageChange={handlePageChange}
                        containerClassName={"flex justify-center mt-4"}
                        pageClassName={
                            "py-1 px-3 cursor-pointer rounded-lg bg-blue-600 text-gray-100 mx-1"
                        }
                        activeClassName={"bg-gray-400 text-white"}
                        previousClassName={
                            "py-1 px-3 cursor-pointer rounded-lg bg-blue-600 text-gray-100"
                        }
                        nextClassName={
                            "py-1 px-3 cursor-pointer rounded-lg bg-blue-600 text-gray-100"
                        }
                        disabledClassName={
                            "py-1 px-3 rounded-lg bg-gray-400 text-gray-100 cursor-not-allowed"
                        }
                    />
                    <div className="flex justify-center mt-6">
                        <button
                            onClick={openModal}
                            className="px-4 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none"
                        >
                            Adicionar Retentor de Valvula
                        </button>
                    </div>
                </div>
                {modalVisible && (
                    <div className="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-75">
                        <div className="bg-white w-1/2 p-6 rounded-lg shadow-lg">
                            <h2 className="text-xl font-semibold mb-4">
                                Novo Retento de Valvula
                            </h2>

                            <div>
                                <input
                                    hidden
                                    readOnly
                                    type="text"
                                    value={id}
                                    className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 w-full"
                                />
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
                                    Codigo:
                                    <input
                                        type="text"
                                        value={codigo}
                                        onChange={handleCodigoChangeModal}
                                        className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 w-full"
                                    />
                                </label>
                                <label className="block mb-2">
                                    Quantidade em Estoque:
                                    <input
                                        type="number"
                                        value={qtd_estoque}
                                        onChange={handleQtdEstoqueChangeModal}
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
                                    onClick={
                                        id == ""
                                            ? handleSave
                                            : () => {
                                                  handleUpdate(id);
                                              }
                                    }
                                    className="px-4 py-2 text-white bg-green-500 rounded-lg hover:bg-green-600 focus:outline-none"
                                >
                                    Salvar
                                </button>
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}
