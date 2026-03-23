import { Link, useNavigate } from 'react-router-dom';
import { clearAuth, getUser } from '../store/auth';

export function Header() {
  const navigate = useNavigate();
  const user = getUser();

  return (
    <header className='header'>
      <div className='container nav'>
        <Link to='/' className='brand'>
          Book Store
        </Link>

        <nav className='nav-links'>
          <Link to='/'>Каталог</Link>
          {user && <Link to='/my-purchases'>Мои покупки</Link>}
          {user && <Link to='/my-rentals'>Мои аренды</Link>}
          {user?.role === 'admin' && <Link to='/admin/books'>Админка</Link>}
        </nav>

        <div className='nav-actions'>
          {user ? (
            <>
              <span>{user.full_name}</span>
              <button
                onClick={() => {
                  clearAuth();
                  navigate('/login');
                }}
              >
                Выйти
              </button>
            </>
          ) : (
            <>
              <Link to='/login'>Войти</Link>
              <Link to='/register'>Регистрация</Link>
            </>
          )}
        </div>
      </div>
    </header>
  );
}
