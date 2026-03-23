import { Navigate } from 'react-router-dom';
import { getUser } from '../store/auth';
import type { JSX } from 'react';

type Props = {
  children: JSX.Element;
  adminOnly?: boolean;
};

export function ProtectedRoute({ children, adminOnly = false }: Props) {
  const user = getUser();

  if (!user) {
    return <Navigate to='/login' replace />;
  }

  if (adminOnly && user.role !== 'admin') {
    return <Navigate to='/' replace />;
  }

  return children;
}
