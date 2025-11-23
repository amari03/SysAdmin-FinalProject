<script>
  import { login } from '../lib/auth.js';

  let email = '';
  let password = '';
  let errorMessage = '';

  async function handleSubmit() {
    errorMessage = ''; // Clear previous errors
    try {
      await login(email, password);
      // The login function handles the redirect on success
    } catch (error) {
      errorMessage = error.message;
    }
  }
</script>

<main>
  <h1>User Login</h1>
  <form on:submit|preventDefault={handleSubmit}>
    <div>
      <label for="email">Email:</label>
      <input type="email" id="email" bind:value={email} required />
    </div>
    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" bind:value={password} required />
    </div>
    <button type="submit">Log In</button>
  </form>

  {#if errorMessage}
    <p style="color: red;">{errorMessage}</p>
  {/if}
</main>

<style>
  main {
    max-width: 400px;
    margin: 2rem auto;
    text-align: center;
  }
  div {
    margin-bottom: 1rem;
  }
  label {
    display: block;
  }
</style>