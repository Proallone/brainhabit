import { writable } from 'svelte/store';

export const token = writable(null);

export function setToken(newToken) {
  token.set(newToken);
}