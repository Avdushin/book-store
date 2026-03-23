import { useEffect, useState } from 'react';
import { getMyPurchases } from '../api/books';
import type { Purchase } from '../types';

export function MyPurchasesPage() {
  const [items, setItems] = useState<Purchase[]>([]);

  useEffect(() => {
    getMyPurchases().then((data) => setItems(data.items));
  }, []);

  return (
    <div>
      <h1>Мои покупки</h1>
      <ul>
        {items.map((item) => (
          <li key={item.id}>
            {item.book_title} — ${item.price}
          </li>
        ))}
      </ul>
    </div>
  );
}
