import React, { useEffect, useState } from "react";
import api from "../../services/api";
import Navbar from "../../components/Navbar";

export default function Junta() {
	const [juntas, setJuntas] = useState([]);
	const [nomePesquisa, setNomePesquisa] = useState("");
	const accessToken = localStorage.getItem("token");

	async function carregarJuntas() {
		const response = await api.get("juntas", {
			headers: {
				Authorization: `Bearer ${accessToken}`,
			},
		});
		setJuntas(response.data);
	}

	const handleNomeChange = (e) => {
		setNomePesquisa(e.target.value);
	};

	const handlePesquisar = () => {
		if (nomePesquisa) {
			api.get(`/juntas/nome/${nomePesquisa}`, {
				headers: {
					Authorization: `Bearer ${accessToken}`,
				},
			})
				.then((response) => {
					const resultado = response.data;
					console.log(resultado);
					setJuntas(resultado);
				})
				.catch((error) => {
					console.error(error);
				});
		} else {
			console.log("Informe um critério de pesquisa válido.");
		}
	};

	async function deleteJunta(id) {
		const confirmDelete = window.confirm("Tem certeza que deseja deletar?");

		if (!confirmDelete) {
			return;
		}
		try {
			await api.delete(`juntas/${id}`, {
				headers: {
					Authorization: `Bearer ${accessToken}`,
				},
			});
			setJuntas(juntas.filter((junta) => junta.id !== id));
		} catch (e) {
			alert("Falha ao deletar junta, tente novamente.");
		}
	}

	useEffect(() => {
		carregarJuntas();
	}, []);

	return (
		<div>
			<Navbar />
			<div className="flex items-center bg-white justify-center h-screen">
				<div className="max-w-md w-full p-6 bg-neutral-500 rounded-lg shadow-md text-center">
					<h1 className="text-2xl font-semibold mb-4">Juntas</h1>

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
								onClick={carregarJuntas}
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
									Qtd Estoque
								</th>
								<th className="py-2 text-center px-4">
									Excluir
								</th>
							</tr>
						</thead>
						<tbody>
							{juntas == null ? (
								<tr>
									<td
										colSpan="6"
										className="py-4 text-center"
									>
										Não há juntas cadastradas.
									</td>
								</tr>
							) : (
								juntas.map((junta) => (
									<tr key={junta.id}>
										<td className="py-2 text-center px-4 border-b">
											{junta.nome}
										</td>
										<td className="py-2 text-center px-4 border-b">
											R$ {junta.preco}
										</td>
										<td className="py-2 text-center px-4 border-b">
											{junta.qtd_estoque}
										</td>
										<button
											className="p-2 text-red-500 hover:text-red-700 focus:outline-none"
											onClick={() =>
												deleteJunta(junta.id)
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
