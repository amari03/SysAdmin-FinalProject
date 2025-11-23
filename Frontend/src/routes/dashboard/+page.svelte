<script>
  import { authStore, logout } from '../../lib/auth.js';

  // The '$' prefix automatically subscribes to the store's value.
  // This means the component will re-render if the auth state changes.
  const user = $authStore.user;
</script>

<main>
  {#if user}
    <h1>Hello {user.name}... Welcome!</h1>
    <p>This is your protected dashboard page.</p>
    <p>Your email is: {user.email}</p>
  {:else}
    <h1>Welcome!</h1>
    <p>Loading user data...</p>
  {/if}

  <button on:click={logout}>Log Out</button>
</main>

<style>
  main {
    padding: 2rem;
    text-align: center;
  }
</style>