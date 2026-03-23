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
  const { data } = await api.get<{ items: Purchase[]; total: number }>('/purchases/my');
  return data;
}

export async function getMyRentals() {
  const { data } = await api.get<{ items: Rental[]; total: number }>('/rentals/my');
  return data;
}

export type AdminCreateBookPayload = {
  title: string;
  description: string;
  author_id: number;
  category_id: number;
  year_written: number;
  purchase_price: number;
  rent_price_2_weeks: number;
  rent_price_1_month: number;
  rent_price_3_months: number;
  status: string;
  is_available: boolean;
  cover_url: string;
};

export async function adminCreateBook(payload: AdminCreateBookPayload) {
  const { data } = await api.post<Book>('/admin/books', payload);
  return data;
}

export async function adminDeleteBook(id: number) {
  const { data } = await api.delete<{ message: string }>(`/admin/books/${id}`);
  return data;
}

export async function adminUpdateBookStatus(id: number, status: string) {
  const { data } = await api.patch<{ message: string }>(`/admin/books/${id}/status`, {
    status,
  });
  return data;
}

export async function adminUpdateBookAvailability(id: number, is_available: boolean) {
  const { data } = await api.patch<{ message: string }>(
    `/admin/books/${id}/availability`,
    {
      is_available,
    },
  );
  return data;
}