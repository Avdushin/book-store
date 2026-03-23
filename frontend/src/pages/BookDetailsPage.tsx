import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { createPurchase, createRental, getBookById } from '../api/books';
import type { Book } from '../types';

export function BookDetailsPage() {
  const { id = '' } = useParams();
  const [book, setBook] = useState<Book | null>(null);
  const [message, setMessage] = useState('');

  async function loadBook() {
    const data = await getBookById(id);
    setBook(data);
  }

  useEffect(() => {
    loadBook();
  }, [id]);

  if (!book) return <p>Загрузка...</p>;

  return (
    <div className='details'>
      <img
        src={`http://localhost:8080${book.cover_url}`}
        alt={book.title}
        className='cover-large'
      />
      <div>
        <h1>{book.title}</h1>
        <p>{book.author_name}</p>
        <p>{book.description}</p>
        <p>Категория: {book.category_name}</p>
        <p>Год: {book.year_written}</p>
        <p>Покупка: ${book.purchase_price}</p>

        <div className='actions'>
          <button
            onClick={async () => {
              try {
                await createPurchase(book.id);
                setMessage('Книга успешно куплена');
                loadBook();
              } catch {
                setMessage('Не удалось купить книгу');
              }
            }}
          >
            Купить
          </button>

          <button
            onClick={async () => {
              try {
                await createRental(book.id, '2_weeks');
                setMessage('Книга арендована на 2 недели');
                loadBook();
              } catch {
                setMessage('Не удалось арендовать книгу');
              }
            }}
          >
            Аренда 2 недели
          </button>

          <button
            onClick={async () => {
              try {
                await createRental(book.id, '1_month');
                setMessage('Книга арендована на 1 месяц');
                loadBook();
              } catch {
                setMessage('Не удалось арендовать книгу');
              }
            }}
          >
            Аренда 1 месяц
          </button>

          <button
            onClick={async () => {
              try {
                await createRental(book.id, '3_months');
                setMessage('Книга арендована на 3 месяца');
                loadBook();
              } catch {
                setMessage('Не удалось арендовать книгу');
              }
            }}
          >
            Аренда 3 месяца
          </button>
        </div>

        {message && <p>{message}</p>}
      </div>
    </div>
  );
}
