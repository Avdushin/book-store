import { api } from './client';
import type { AuthResponse, User } from '../types';

export async function login(email: string, password: string) {
  const { data } = await api.post<AuthResponse>('/auth/login', {
    email,
    password,
  });
  return data;
}

export async function register(
  full_name: string,
  email: string,
  password: string,
) {
  const { data } = await api.post<AuthResponse>('/auth/register', {
    full_name,
    email,
    password,
  });
  return data;
}

export async function me() {
  const { data } = await api.get<User>('/auth/me');
  return data;
}
