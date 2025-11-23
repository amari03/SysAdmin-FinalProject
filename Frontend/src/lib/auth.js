import { writable } from 'svelte/store';
import { goto } from '$app/navigation';

// The store will hold the user's login status and their data.
const initialAuthState = { loggedIn: false, user: null, token: null };
export const authStore = writable(initialAuthState);

// --- Helper function to log the user in ---
export const login = async (email, password) => {
  try {
    // We will call your Go backend's authentication endpoint.
    const response = await fetch('http://localhost:4000/v1/tokens/authentication', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
      // If the server returns an error (like 401 Unauthorized), throw an error.
      throw new Error('Invalid email or password.');
    }

    const data = await response.json();

    // On success, update the store with the user data and token.
    authStore.set({
      loggedIn: true,
      user: data.user,
      token: data.authentication_token.token,
    });

    // Redirect the user to their dashboard.
    goto('/dashboard');

  } catch (error) {
    // If anything fails, clear the store and re-throw the error
    // so the login page can display a message.
    authStore.set(initialAuthState);
    throw error;
  }
};

// --- Helper function to log the user out ---
export const logout = () => {
  authStore.set(initialAuthState);
  // Redirect to the login page.
  goto('/');
};