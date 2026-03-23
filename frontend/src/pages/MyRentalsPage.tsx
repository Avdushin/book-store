import { useEffect, useState } from 'react';
import { getMyRentals } from '../api/books';
import type { Rental } from '../types';

export function MyRentalsPage() {
  const [items, setItems] = useState<Rental[]>([]);

  useEffect(() => {
    getMyRentals().then((data) => setItems(data.items));
  }, []);

  return (
    <div>
      <h1>Мои аренды</h1>
      <ul>
        {items.map((item) => (
          <li key={item.id}>
            {item.book_title} — {item.tariff} — {item.status}
          </li>
        ))}
      </ul>
    </div>
  );
}
