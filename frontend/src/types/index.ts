export type User = {
  id: number;
  full_name: string;
  email: string;
  role: 'user' | 'admin';
};

export type AuthResponse = {
  token: string;
  user: User;
};

export type Book = {
  id: number;
  title: string;
  description: string;
  author_id: number;
  author_name: string;
  category_id: number;
  category_name: string;
  year_written: number;
  purchase_price: number;
  rent_price_2_weeks: number;
  rent_price_1_month: number;
  rent_price_3_months: number;
  status: string;
  is_available: boolean;
  cover_url: string;
};

export type BookListResponse = {
  items: Book[];
  total: number;
};

export type Purchase = {
  id: number;
  user_id: number;
  book_id: number;
  book_title: string;
  price: number;
  purchased_at: string;
};

export type Rental = {
  id: number;
  user_id: number;
  book_id: number;
  book_title: string;
  tariff: string;
  start_date: string;
  end_date: string;
  status: string;
  created_at: string;
};
