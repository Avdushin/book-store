import { api } from './client';
import type { Book, BookListResponse, Purchase, Rental } from '../types';

export async function getBooks(params?: Record<string, string | number>) {
  const { data } = await api.get<BookListResponse>('/books', { params });
  return data;
}

export async function getBookById(id: string) {
  const { data } = await api.get<Book>(`/books/${id}`);
  return data;
}

export async function createPurchase(book_id: number) {
  const { data } = await api.post<Purchase>('/purchases', { book_id });
  return data;
}

export async function createRental(book_id: number, tariff: string) {
  const { data } = await api.post<Rental>('/rentals', { book_id, tariff });
  return data;
}

export async function getMyPurchases() {
  const { data } = await api.get<{ items: Purchase[]; total: number }>(
    '/purchases/my',
  );
  return data;
}

export async function getMyRentals() {
  const { data } = await api.get<{ items: Rental[]; total: number }>(
    '/rentals/my',
  );
  return data;
}
