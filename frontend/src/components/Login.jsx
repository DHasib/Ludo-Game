import React, { useState } from 'react';
import { api } from '../api';

export default function Login({ onLogin }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  async function handleLogin(e) {
    e.preventDefault();
    const res = await api.post('/api/auth/login', { email, password });
    onLogin(res.token);
  }

  async function handleRegister(e) {
    e.preventDefault();
    await api.post('/api/auth/register', { email, password });
    await handleLogin(e);
  }

  return (
    <form>
      <input value={email} onChange={e => setEmail(e.target.value)} placeholder="email" />
      <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="password" />
      <button onClick={handleLogin}>Login</button>
      <button onClick={handleRegister}>Register</button>
    </form>
  );
}
