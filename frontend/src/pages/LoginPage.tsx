import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { login } from '../api/auth';
import { setAuth } from '../store/auth';

export function LoginPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState('user@bookstore.com');
  const [password, setPassword] = useState('user123');
  const [error, setError] = useState('');

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError('');

    try {
      const data = await login(email, password);
      setAuth(data.token, data.user);
      navigate(data.user.role === 'admin' ? '/admin/books' : '/');
    } catch {
      setError('Не удалось войти');
    }
  }

  return (
    <div className='auth-box'>
      <h1>Вход</h1>

      <form onSubmit={onSubmit} className='auth-form'>
        <input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder='Email'
          type='email'
        />
        <input
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder='Пароль'
          type='password'
        />
        <button type='submit'>Войти</button>
      </form>

      {error && <p className='auth-error'>{error}</p>}

      <p className='auth-switch'>
        Нет аккаунта? <Link to='/register'>Зарегистрироваться</Link>
      </p>
    </div>
  );
}
