<script>
  import { goto } from '$app/navigation';

  let name = '';
  let email = '';
  let password = '';
  let message = '';
  let isError = false;

  async function handleSignUp() {
    message = '';
    isError = false;
    try {
      const response = await fetch('http://localhost:4000/v1/users', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password }),
      });
      const data = await response.json();
      if (!response.ok) {
        const errorMessage = data.error?.message || JSON.stringify(data.error);
        throw new Error(errorMessage || 'Sign up failed.');
      }
      message = "Success! Please check your email to activate your account.";
    } catch (error) {
      isError = true;
      message = error.message;
    }
  }

  function goToLogin() {
    goto('/');
  }
</script>

<main>
  <div class="signup-container">
    <div class="header">
      <div class="icon">
        <div class="emoji-carousel">
          <span>😊</span>
          <span>😔</span>
          <span>😌</span>
          <span>🥰</span>
          <span>😤</span>
          <span>😴</span>
          <span>🤗</span>
          <span>😢</span>
          <!-- Duplicate for seamless loop -->
          <span>😊</span>
          <span>😔</span>
          <span>😌</span>
          <span>🥰</span>
          <span>😤</span>
          <span>😴</span>
          <span>🤗</span>
          <span>😢</span>
        </div>
      </div>
      <h1>Join Feel Flow</h1>
      <p class="subtitle">Start tracking your emotional journey</p>
    </div>

    <form on:submit|preventDefault={handleSignUp}>
      <div class="input-group">
        <label for="name">Name</label>
        <input 
          type="text" 
          id="name" 
          bind:value={name} 
          placeholder="Your name"
          required 
        />
      </div>

      <div class="input-group">
        <label for="email">Email</label>
        <input 
          type="email" 
          id="email" 
          bind:value={email} 
          placeholder="you@example.com"
          required 
        />
      </div>

      <div class="input-group">
        <label for="password">Password</label>
        <input 
          type="password" 
          id="password" 
          bind:value={password} 
          placeholder="••••••••"
          required 
        />
      </div>

      {#if message && isError}
        <div class="error-message">
          <span class="error-icon">⚠️</span>
          {message}
        </div>
      {/if}

      {#if message && !isError}
        <div class="success-message">
          <span class="success-icon">✓</span>
          {message}
        </div>
      {/if}

      <button type="submit" class="primary-button">Create Account</button>
      
      <div class="divider">
        <span>or</span>
      </div>

      <button type="button" class="secondary-button" on:click={goToLogin}>
        Back to Login
      </button>
    </form>

    <p class="terms">
      By signing up, you agree to our Terms of Service and Privacy Policy
    </p>
  </div>
</main>

<style>
  :global(body) {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  }

  main {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 2rem 1rem;
  }

  .signup-container {
    background: white;
    border-radius: 24px;
    padding: 3rem 2.5rem;
    max-width: 420px;
    width: 100%;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  }

  .header {
    text-align: center;
    margin-bottom: 2.5rem;
  }

  .icon {
    font-size: 3rem;
    margin-bottom: 1rem;
    height: 3.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    position: relative;
  }

  .emoji-carousel {
    display: flex;
    animation: scroll 12s linear infinite;
  }

  .emoji-carousel span {
    min-width: 5rem;
    text-align: center;
  }

  @keyframes scroll {
    0% {
      transform: translateX(0);
    }
    100% {
      transform: translateX(calc(-5rem * 8));
    }
  }

  h1 {
    font-size: 2rem;
    font-weight: 700;
    color: #1a202c;
    margin: 0 0 0.5rem 0;
  }

  .subtitle {
    color: #718096;
    font-size: 0.95rem;
    margin: 0;
  }

  form {
    display: flex;
    flex-direction: column;
  }

  .input-group {
    margin-bottom: 1.5rem;
    text-align: left;
  }

  label {
    display: block;
    font-size: 0.875rem;
    font-weight: 600;
    color: #4a5568;
    margin-bottom: 0.5rem;
  }

  input {
    width: 100%;
    padding: 0.875rem 1rem;
    border: 2px solid #e2e8f0;
    border-radius: 12px;
    font-size: 1rem;
    transition: all 0.2s ease;
    box-sizing: border-box;
  }

  input:focus {
    outline: none;
    border-color: #667eea;
    box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
  }

  input::placeholder {
    color: #cbd5e0;
  }

  .primary-button {
    width: 100%;
    padding: 1rem;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 12px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
  }

  .primary-button:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(102, 126, 234, 0.5);
  }

  .primary-button:active {
    transform: translateY(0);
  }

  .divider {
    display: flex;
    align-items: center;
    margin: 1.5rem 0;
    color: #a0aec0;
    font-size: 0.875rem;
  }

  .divider::before,
  .divider::after {
    content: '';
    flex: 1;
    height: 1px;
    background: #e2e8f0;
  }

  .divider span {
    padding: 0 1rem;
  }

  .secondary-button {
    width: 100%;
    padding: 1rem;
    background: white;
    color: #667eea;
    border: 2px solid #667eea;
    border-radius: 12px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .secondary-button:hover {
    background: #f7fafc;
    transform: translateY(-1px);
  }

  .error-message {
    background: #fff5f5;
    color: #c53030;
    padding: 0.875rem 1rem;
    border-radius: 12px;
    margin-bottom: 1.5rem;
    font-size: 0.875rem;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    border: 1px solid #feb2b2;
  }

  .error-icon {
    font-size: 1.1rem;
  }

  .success-message {
    background: #f0fff4;
    color: #276749;
    padding: 0.875rem 1rem;
    border-radius: 12px;
    margin-bottom: 1.5rem;
    font-size: 0.875rem;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    border: 1px solid #9ae6b4;
  }

  .success-icon {
    font-size: 1.1rem;
    font-weight: bold;
  }

  .terms {
    text-align: center;
    font-size: 0.75rem;
    color: #a0aec0;
    margin-top: 1.5rem;
    line-height: 1.4;
  }

  @media (max-width: 480px) {
    .signup-container {
      padding: 2rem 1.5rem;
    }

    h1 {
      font-size: 1.75rem;
    }
  }
</style>