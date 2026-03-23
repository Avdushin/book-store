import { useEffect, useState } from 'react';
import { getBooks } from '../api/books';
import { api } from '../api/client';
import type { Book } from '../types';

export function AdminBooksPage() {
  const [items, setItems] = useState<Book[]>([]);
  const [title, setTitle] = useState('');

  async function loadBooks() {
    const data = await getBooks();
    setItems(data.items);
  }

  useEffect(() => {
    loadBooks();
  }, []);

  async function createBook() {
    await api.post('/admin/books', {
      title,
      description: 'Новая книга',
      author_id: 1,
      category_id: 1,
      year_written: 2024,
      purchase_price: 10,
      rent_price_2_weeks: 2,
      rent_price_1_month: 4,
      rent_price_3_months: 8,
      status: 'available',
      is_available: true,
      cover_url: '/covers/test.jpg',
    });

    setTitle('');
    loadBooks();
  }

  return (
    <div>
      <h1>Админка книг</h1>

      <div className='filters'>
        <input
          placeholder='Название новой книги'
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <button onClick={createBook}>Добавить</button>
      </div>

      <ul>
        {items.map((book) => (
          <li key={book.id}>
            {book.title} — {book.status} — {String(book.is_available)}
          </li>
        ))}
      </ul>
    </div>
  );
}
