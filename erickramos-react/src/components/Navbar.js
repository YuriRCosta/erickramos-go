import React, { useState } from "react";
import { Link, NavLink } from "react-router-dom";

const Navbar = () => {
	const [dropdownOpen, setDropdownOpen] = useState(false);

	const toggleDropdown = () => {
		setDropdownOpen(!dropdownOpen);
	};

	return (
		<nav className="bg-gray-800 py-4 absolute w-full">
			<div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div className="flex justify-between">
					<div>
						<Link href="#" className="text-white font-bold text-xl">
							Erick Ramos
						</Link>
					</div>
					<div className="flex">
						<div
							className="relative inline-block text-left mr-2"
							role="menu"
						>
							<button
								type="button"
								className="inline-flex justify-center w-full rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-gray-700 text-sm font-medium text-white hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-100 focus:ring-indigo-500"
							>
								<Link href="/orcamento">
									<div className="text-white font-medium hover:text-gray-300">
										Orçamento
									</div>
								</Link>
							</button>
						</div>
						<div className="relative inline-block text-left">
							<button
								type="button"
								onClick={toggleDropdown}
								className="inline-flex justify-center w-full rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-gray-700 text-sm font-medium text-white hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-100 focus:ring-indigo-500"
								aria-expanded={dropdownOpen}
							>
								Estoque
								<svg
									className="-mr-1 ml-2 h-5 w-5"
									xmlns="http://www.w3.org/2000/svg"
									viewBox="0 0 20 20"
									fill="currentColor"
									aria-hidden="true"
								>
									<path
										fillRule="evenodd"
										d="M10 3a1 1 0 00-.707.293l-6 6a1 1 0 000 1.414l6 6A1 1 0 0010 16H7a3 3 0 01-3-3V6a3 3 0 013-3h3zm5.293 7.707a1 1 0 010 1.414l-6 6A1 1 0 018 16H5a1 1 0 01-1-1V6a1 1 0 011-1h3a1 1 0 01.707.293l6 6z"
										clipRule="evenodd"
									/>
								</svg>
							</button>
							{dropdownOpen && (
								<div className="origin-top-right absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5">
									<div
										className="py-1"
										role="menu"
										aria-orientation="vertical"
										aria-labelledby="options-menu"
									>
										<Link
											to="/retentoresComando"
											className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"
											role="menuitem"
										>
											Retentor de Comando
										</Link>
										<Link
											to="/retentoresValvula"
											className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"
											role="menuitem"
										>
											Retentor de Valvula
										</Link>
									</div>
								</div>
							)}
						</div>
					</div>
				</div>
			</div>
		</nav>
	);
};

export default Navbar;
