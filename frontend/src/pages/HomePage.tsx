import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { getBooks } from '../api/books';
import type { Book } from '../types';

export function HomePage() {
  const [books, setBooks] = useState<Book[]>([]);
  const [category, setCategory] = useState('');
  const [author, setAuthor] = useState('');
  const [year, setYear] = useState('');
  const [sortBy, setSortBy] = useState('title');
  const [order, setOrder] = useState('asc');

  async function loadBooks() {
    const data = await getBooks({
      ...(category ? { category } : {}),
      ...(author ? { author } : {}),
      ...(year ? { year } : {}),
      sort_by: sortBy,
      order,
    });
    setBooks(data.items);
  }

  useEffect(() => {
    loadBooks();
  }, []);

  return (
    <div>
      <h1>Каталог книг</h1>

      <div className='filters'>
        <input
          placeholder='Категория'
          value={category}
          onChange={(e) => setCategory(e.target.value)}
        />
        <input
          placeholder='Автор'
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
        />
        <input
          placeholder='Год'
          value={year}
          onChange={(e) => setYear(e.target.value)}
        />
        <select value={sortBy} onChange={(e) => setSortBy(e.target.value)}>
          <option value='title'>Название</option>
          <option value='author'>Автор</option>
          <option value='category'>Категория</option>
          <option value='year'>Год</option>
          <option value='price'>Цена</option>
        </select>
        <select value={order} onChange={(e) => setOrder(e.target.value)}>
          <option value='asc'>ASC</option>
          <option value='desc'>DESC</option>
        </select>
        <button onClick={loadBooks}>Применить</button>
      </div>

      <div className='grid'>
        {books.map((book) => (
          <article key={book.id} className='card'>
            <img
              src={`http://localhost:8080${book.cover_url}`}
              alt={book.title}
              className='cover'
            />
            <h3>{book.title}</h3>
            <p>{book.author_name}</p>
            <p>{book.category_name}</p>
            <p>{book.year_written}</p>
            <Link to={`/books/${book.id}`}>Открыть</Link>
          </article>
        ))}
      </div>
    </div>
  );
}
