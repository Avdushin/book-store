import { Outlet } from 'react-router-dom';
import { Header } from '../components/Header';

export function App() {
  return (
    <div className='app-shell'>
      <Header />
      <main className='container'>
        <Outlet />
      </main>
    </div>
  );
}
