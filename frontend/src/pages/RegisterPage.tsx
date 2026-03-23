import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { register } from '../api/auth';
import { setAuth } from '../store/auth';

export function RegisterPage() {
  const navigate = useNavigate();

  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError('');

    try {
      const data = await register(fullName, email, password);
      setAuth(data.token, data.user);
      navigate('/');
    } catch {
      setError('Не удалось зарегистрироваться');
    }
  }

  return (
    <div className='auth-box'>
      <h1>Регистрация</h1>

      <form onSubmit={onSubmit} className='auth-form'>
        <input
          value={fullName}
          onChange={(e) => setFullName(e.target.value)}
          placeholder='Полное имя'
        />
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
        <button type='submit'>Зарегистрироваться</button>
      </form>

      {error && <p className='auth-error'>{error}</p>}

      <p className='auth-switch'>
        Уже есть аккаунт? <Link to='/login'>Войти</Link>
      </p>
    </div>
  );
}
