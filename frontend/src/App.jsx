import React, { useState } from 'react';
import Login from './components/Login';
import GameBoard from './components/GameBoard';

export default function App() {
  const [token, setToken] = useState(localStorage.getItem('token'));

  if (!token) {
    return <Login onLogin={t => { localStorage.setItem('token', t); setToken(t); }} />;
  }

  return <GameBoard token={token} />;
}
