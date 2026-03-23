import { useEffect, useState } from 'react';
import axios from 'axios';
import {
  adminCreateBook,
  adminDeleteBook,
  adminUpdateBookAvailability,
  adminUpdateBookStatus,
  getAuthors,
  getBooks,
  getCategories,
} from '../api/books';
import type { Author, Book, Category } from '../types';

type CreateFormState = {
  title: string;
  description: string;
  author_id: string;
  category_id: string;
  year_written: string;
  purchase_price: string;
  rent_price_2_weeks: string;
  rent_price_1_month: string;
  rent_price_3_months: string;
  status: string;
  is_available: boolean;
  cover_url: string;
};

const initialForm: CreateFormState = {
  title: '',
  description: '',
  author_id: '',
  category_id: '',
  year_written: '2024',
  purchase_price: '10',
  rent_price_2_weeks: '2',
  rent_price_1_month: '4',
  rent_price_3_months: '8',
  status: 'available',
  is_available: true,
  cover_url: '/covers/test.jpg',
};

function getStatusLabel(book: Book) {
  if (book.status === 'sold_out') {
    return 'Продана';
  }

  if (book.status === 'rented') {
    return 'Арендована';
  }

  if (book.status === 'inactive') {
    return 'Скрыта';
  }

  if (book.status === 'available' && book.is_available) {
    return 'Доступна';
  }

  return 'Недоступна';
}

export function AdminBooksPage() {
  const [items, setItems] = useState<Book[]>([]);
  const [authors, setAuthors] = useState<Author[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [form, setForm] = useState<CreateFormState>(initialForm);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  async function loadBooks() {
    const data = await getBooks({ sort_by: 'title', order: 'asc' });
    setItems(data.items);
  }

  async function loadReferences() {
    const [authorsData, categoriesData] = await Promise.all([
      getAuthors(),
      getCategories(),
    ]);

    setAuthors(authorsData.items);
    setCategories(categoriesData.items);

    setForm((prev) => ({
      ...prev,
      author_id: prev.author_id || String(authorsData.items[0]?.id ?? ''),
      category_id:
        prev.category_id || String(categoriesData.items[0]?.id ?? ''),
    }));
  }

  async function loadAll() {
    try {
      setError('');
      await Promise.all([loadBooks(), loadReferences()]);
    } catch {
      setError('Не удалось загрузить данные админки');
    }
  }

  useEffect(() => {
    loadAll();
  }, []);

  function updateForm<K extends keyof CreateFormState>(
    key: K,
    value: CreateFormState[K],
  ) {
    setForm((prev) => ({
      ...prev,
      [key]: value,
    }));
  }

  async function handleCreateBook(e: React.FormEvent) {
    e.preventDefault();
    setError('');
    setMessage('');
    setLoading(true);

    try {
      await adminCreateBook({
        title: form.title,
        description: form.description,
        author_id: Number(form.author_id),
        category_id: Number(form.category_id),
        year_written: Number(form.year_written),
        purchase_price: Number(form.purchase_price),
        rent_price_2_weeks: Number(form.rent_price_2_weeks),
        rent_price_1_month: Number(form.rent_price_1_month),
        rent_price_3_months: Number(form.rent_price_3_months),
        status: form.status,
        is_available: form.is_available,
        cover_url: form.cover_url,
      });

      setMessage('Книга успешно добавлена');
      setForm((prev) => ({
        ...initialForm,
        author_id: prev.author_id,
        category_id: prev.category_id,
      }));

      await loadBooks();
    } catch (err) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.error ?? 'Не удалось создать книгу');
      } else {
        setError('Не удалось создать книгу');
      }
    } finally {
      setLoading(false);
    }
  }

  async function handleDeleteBook(book: Book) {
    const confirmed = window.confirm(`Удалить книгу "${book.title}"?`);
    if (!confirmed) {
      return;
    }

    setError('');
    setMessage('');

    try {
      await adminDeleteBook(book.id);
      setMessage(`Книга "${book.title}" удалена`);
      await loadBooks();
    } catch (err) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.error ?? 'Не удалось удалить книгу');
      } else {
        setError('Не удалось удалить книгу');
      }
    }
  }

  async function handleStatusChange(book: Book, status: string) {
    setError('');
    setMessage('');

    try {
      await adminUpdateBookStatus(book.id, status);
      setMessage(`Статус книги "${book.title}" обновлён`);
      await loadBooks();
    } catch (err) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.error ?? 'Не удалось обновить статус');
      } else {
        setError('Не удалось обновить статус');
      }
    }
  }

  async function handleAvailabilityToggle(book: Book) {
    setError('');
    setMessage('');

    try {
      await adminUpdateBookAvailability(book.id, !book.is_available);
      setMessage(`Доступность книги "${book.title}" обновлена`);
      await loadBooks();
    } catch (err) {
      if (axios.isAxiosError(err)) {
        setError(
          err.response?.data?.error ?? 'Не удалось обновить доступность',
        );
      } else {
        setError('Не удалось обновить доступность');
      }
    }
  }

  return (
    <div className='admin-page'>
      <h1>Админка книг</h1>

      <section className='admin-section'>
        <h2>Добавить новую книгу</h2>

        <form className='admin-form' onSubmit={handleCreateBook}>
          <input
            placeholder='Название'
            value={form.title}
            onChange={(e) => updateForm('title', e.target.value)}
          />

          <input
            placeholder='Описание'
            value={form.description}
            onChange={(e) => updateForm('description', e.target.value)}
          />

          <select
            value={form.author_id}
            onChange={(e) => updateForm('author_id', e.target.value)}
          >
            <option value=''>Выберите автора</option>
            {authors.map((author) => (
              <option key={author.id} value={author.id}>
                {author.full_name}
              </option>
            ))}
          </select>

          <select
            value={form.category_id}
            onChange={(e) => updateForm('category_id', e.target.value)}
          >
            <option value=''>Выберите категорию</option>
            {categories.map((category) => (
              <option key={category.id} value={category.id}>
                {category.name}
              </option>
            ))}
          </select>

          <input
            placeholder='Год'
            value={form.year_written}
            onChange={(e) => updateForm('year_written', e.target.value)}
          />

          <input
            placeholder='Цена покупки'
            value={form.purchase_price}
            onChange={(e) => updateForm('purchase_price', e.target.value)}
          />

          <input
            placeholder='Аренда 2 недели'
            value={form.rent_price_2_weeks}
            onChange={(e) => updateForm('rent_price_2_weeks', e.target.value)}
          />

          <input
            placeholder='Аренда 1 месяц'
            value={form.rent_price_1_month}
            onChange={(e) => updateForm('rent_price_1_month', e.target.value)}
          />

          <input
            placeholder='Аренда 3 месяца'
            value={form.rent_price_3_months}
            onChange={(e) => updateForm('rent_price_3_months', e.target.value)}
          />

          <input
            placeholder='Путь к обложке'
            value={form.cover_url}
            onChange={(e) => updateForm('cover_url', e.target.value)}
          />

          <select
            value={form.status}
            onChange={(e) => updateForm('status', e.target.value)}
          >
            <option value='available'>available</option>
            <option value='rented'>rented</option>
            <option value='sold_out'>sold_out</option>
            <option value='inactive'>inactive</option>
          </select>

          <label className='checkbox-row'>
            <input
              type='checkbox'
              checked={form.is_available}
              onChange={(e) => updateForm('is_available', e.target.checked)}
            />
            Доступна
          </label>

          <button type='submit' disabled={loading}>
            {loading ? 'Сохраняем...' : 'Добавить книгу'}
          </button>
        </form>

        {message && <p className='admin-message success'>{message}</p>}
        {error && <p className='admin-message error'>{error}</p>}
      </section>

      <section className='admin-section'>
        <h2>Список книг</h2>

        <div className='admin-books-grid'>
          {items.map((book) => (
            <article key={book.id} className='admin-book-card'>
              <img
                src={`http://localhost:8080${book.cover_url}`}
                alt={book.title}
                className='admin-book-cover'
              />

              <div className='admin-book-content'>
                <h3>{book.title}</h3>

                <p>
                  <strong>ID:</strong> {book.id}
                </p>

                <p>
                  <strong>Автор:</strong> {book.author_name}
                </p>

                <p>
                  <strong>Категория:</strong> {book.category_name}
                </p>

                <p>
                  <strong>Год:</strong> {book.year_written}
                </p>

                <p>
                  <strong>Цена:</strong> ${book.purchase_price}
                </p>

                <p>
                  <strong>Статус:</strong> {getStatusLabel(book)}
                </p>

                <p>
                  <strong>Служебный статус:</strong> {book.status}
                </p>

                <p>
                  <strong>is_available:</strong> {String(book.is_available)}
                </p>

                <div className='admin-actions'>
                  <select
                    value={book.status}
                    onChange={(e) => handleStatusChange(book, e.target.value)}
                  >
                    <option value='available'>available</option>
                    <option value='rented'>rented</option>
                    <option value='sold_out'>sold_out</option>
                    <option value='inactive'>inactive</option>
                  </select>

                  <button onClick={() => handleAvailabilityToggle(book)}>
                    {book.is_available
                      ? 'Сделать недоступной'
                      : 'Сделать доступной'}
                  </button>

                  <button
                    className='danger-button'
                    onClick={() => handleDeleteBook(book)}
                  >
                    Удалить
                  </button>
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>
    </div>
  );
}
