import { createBrowserRouter } from 'react-router-dom';
import { App } from '../app/App';
import { ProtectedRoute } from '../components/ProtectedRoute';
import { AdminBooksPage } from '../pages/AdminBooksPage';
import { BookDetailsPage } from '../pages/BookDetailsPage';
import { HomePage } from '../pages/HomePage';
import { LoginPage } from '../pages/LoginPage';
import { MyPurchasesPage } from '../pages/MyPurchasesPage';
import { MyRentalsPage } from '../pages/MyRentalsPage';
import { NotFoundPage } from '../pages/NotFoundPage';
import { RegisterPage } from '../pages/RegisterPage';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      { index: true, element: <HomePage /> },
      { path: '/register', element: <RegisterPage /> },
      { path: 'login', element: <LoginPage /> },
      { path: 'books/:id', element: <BookDetailsPage /> },
      {
        path: 'my-purchases',
        element: (
          <ProtectedRoute>
            <MyPurchasesPage />
          </ProtectedRoute>
        ),
      },
      {
        path: 'my-rentals',
        element: (
          <ProtectedRoute>
            <MyRentalsPage />
          </ProtectedRoute>
        ),
      },
      {
        path: 'admin/books',
        element: (
          <ProtectedRoute adminOnly>
            <AdminBooksPage />
          </ProtectedRoute>
        ),
      },
      { path: '*', element: <NotFoundPage /> },
    ],
  },
]);
