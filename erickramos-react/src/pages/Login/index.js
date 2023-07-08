import React, { useState } from 'react';
import Api from "../../services/api";

export default function Login({children}) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  async function handleLogin(e) {
    e.preventDefault();

    const data = {
        username,
        password
    };

    try {
        const response = await Api.post('/login', data);
        localStorage.setItem('token', response.data.token);
        console.log(response.data);
    } catch (e) {
        alert('Falha no login, tente novamente.');
    }
}

  return (
    <div className="flex justify-center items-center h-screen">
      <form onSubmit={handleLogin} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <h1 className="text-2xl mb-6">Login</h1>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2">Username:</label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            type="username"
            value={username}
            onChange={e => setUsername(e.target.value)}
          />
        </div>
        <div className="mb-6">
          <label className="block text-gray-700 text-sm font-bold mb-2">Password:</label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            type="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
          />
        </div>
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          type="submit"
        >
          Login
        </button>
      </form>
    </div>
  );
};

