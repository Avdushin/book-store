import { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { createPurchase, createRental, getBookById } from '../api/books';
import type { Book } from '../types';

function getAvailabilityText(book: Book) {
  if (book.is_available && book.status === 'available') {
    return 'Доступна';
  }

  if (book.status === 'sold_out') {
    return 'Продана';
  }

  if (book.status === 'rented') {
    return 'Сейчас арендована';
  }

  if (book.status === 'inactive') {
    return 'Скрыта администратором';
  }

  return 'Недоступна';
}

function isBookActionAvailable(book: Book) {
  return book.is_available && book.status === 'available';
}

export function BookDetailsPage() {
  const { id = '' } = useParams();
  const [book, setBook] = useState<Book | null>(null);
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(true);

  async function loadBook() {
    try {
      setLoading(true);
      const data = await getBookById(id);
      setBook(data);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        setMessage(error.response?.data?.error ?? 'Не удалось загрузить книгу');
      } else {
        setMessage('Не удалось загрузить книгу');
      }
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadBook();
  }, [id]);

  async function handlePurchase() {
    if (!book) return;

    setMessage('');

    try {
      await createPurchase(book.id);
      setMessage('Книга успешно куплена');
      await loadBook();
    } catch (error) {
      if (axios.isAxiosError(error)) {
        setMessage(error.response?.data?.error ?? 'Не удалось купить книгу');
      } else {
        setMessage('Не удалось купить книгу');
      }
    }
  }

  async function handleRental(tariff: '2_weeks' | '1_month' | '3_months') {
    if (!book) return;

    setMessage('');

    try {
      await createRental(book.id, tariff);

      if (tariff === '2_weeks') {
        setMessage('Книга успешно арендована на 2 недели');
      } else if (tariff === '1_month') {
        setMessage('Книга успешно арендована на 1 месяц');
      } else {
        setMessage('Книга успешно арендована на 3 месяца');
      }

      await loadBook();
    } catch (error) {
      if (axios.isAxiosError(error)) {
        setMessage(error.response?.data?.error ?? 'Не удалось арендовать книгу');
      } else {
        setMessage('Не удалось арендовать книгу');
      }
    }
  }

  if (loading) {
    return <p>Загрузка...</p>;
  }

  if (!book) {
    return <p>Книга не найдена</p>;
  }

  const available = isBookActionAvailable(book);

  return (
    <div className='details'>
      <img
        src={`http://localhost:8080${book.cover_url}`}
        alt={book.title}
        className='cover-large'
      />

      <div className='details-content'>
        <h1>{book.title}</h1>

        <p>
          <strong>Автор:</strong> {book.author_name}
        </p>

        <p>
          <strong>Категория:</strong> {book.category_name}
        </p>

        <p>
          <strong>Год написания:</strong> {book.year_written}
        </p>

        <p>
          <strong>Статус:</strong> {book.status}
        </p>

        <p>
          <strong>Доступность:</strong> {getAvailabilityText(book)}
        </p>

        <p>
          <strong>Описание:</strong> {book.description}
        </p>

        <p>
          <strong>Цена покупки:</strong> ${book.purchase_price}
        </p>

        <p>
          <strong>Аренда на 2 недели:</strong> ${book.rent_price_2_weeks}
        </p>

        <p>
          <strong>Аренда на 1 месяц:</strong> ${book.rent_price_1_month}
        </p>

        <p>
          <strong>Аренда на 3 месяца:</strong> ${book.rent_price_3_months}
        </p>

        <div className='actions'>
          <button onClick={handlePurchase} disabled={!available}>
            Купить
          </button>

          <button onClick={() => handleRental('2_weeks')} disabled={!available}>
            Аренда 2 недели
          </button>

          <button onClick={() => handleRental('1_month')} disabled={!available}>
            Аренда 1 месяц
          </button>

          <button onClick={() => handleRental('3_months')} disabled={!available}>
            Аренда 3 месяца
          </button>
        </div>

        {!available && (
          <p className='status-note'>
            Эта книга сейчас недоступна для покупки и аренды.
          </p>
        )}

        {message && <p className='details-message'>{message}</p>}
      </div>
    </div>
  );
}