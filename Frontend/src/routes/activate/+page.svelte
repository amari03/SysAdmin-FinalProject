<script>
  import { onMount } from 'svelte';
  import { page } from '$app/stores';

  let message = 'Activating your account, please wait...';
  let isError = false;
  let isLoading = true;

  // onMount runs once after the component is first rendered in the browser.
  onMount(async () => {
    // Get the token from the URL's query parameters.
    const token = $page.url.searchParams.get('token');

    if (!token) {
      isError = true;
      isLoading = false;
      message = 'Activation token not found in URL. Please check the link in your email.';
      return;
    }

    try {
      const response = await fetch('http://localhost:4000/v1/users/activated', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ token: token }),
      });

      if (!response.ok) {
        throw new Error('Activation failed. The token may be invalid or expired.');
      }

      // If we get here, activation was successful!
      isError = false;
      isLoading = false;
      message = 'Success! Your account has been activated. You can now log in.';

    } catch (error) {
      isError = true;
      isLoading = false;
      message = error.message;
    }
  });
</script>

<div class="page-wrapper">
  <main>
    <div class="card">
      <div class="icon-wrapper">
        {#if isLoading}
          <div class="spinner"></div>
        {:else if isError}
          <span class="icon error-icon">❌</span>
        {:else}
          <span class="icon success-icon">✅</span>
        {/if}
      </div>

      <h1>Account Activation</h1>
      
      <p class="message" class:error={isError} class:success={!isError && !isLoading}>
        {message}
      </p>

      {#if !isError && message.startsWith('Success')}
        <a href="/" class="login-link">
          <span>Go to Login Page</span>
          <span class="arrow">→</span>
        </a>
      {/if}
    </div>
  </main>
</div>

<style>
  :global(body) {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    margin: 0;
    padding: 0;
  }

  .page-wrapper {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem;
  }

  main {
    width: 100%;
    max-width: 500px;
  }

  .card {
    background: white;
    border-radius: 20px;
    padding: 3rem 2.5rem;
    text-align: center;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
    animation: slideUp 0.5s ease;
  }

  @keyframes slideUp {
    from {
      transform: translateY(30px);
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }

  .icon-wrapper {
    margin-bottom: 1.5rem;
    height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .icon {
    font-size: 4rem;
    display: inline-block;
  }

  .success-icon {
    animation: successBounce 0.6s ease;
  }

  @keyframes successBounce {
    0%, 100% {
      transform: scale(1);
    }
    50% {
      transform: scale(1.2);
    }
  }

  .error-icon {
    animation: shake 0.5s ease;
  }

  @keyframes shake {
    0%, 100% {
      transform: translateX(0);
    }
    25% {
      transform: translateX(-10px);
    }
    75% {
      transform: translateX(10px);
    }
  }

  .spinner {
    width: 60px;
    height: 60px;
    border: 5px solid #e2e8f0;
    border-top-color: #667eea;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  h1 {
    margin: 0 0 1.5rem 0;
    font-size: 2rem;
    font-weight: 700;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  .message {
    margin: 0 0 2rem 0;
    font-size: 1.1rem;
    line-height: 1.6;
    color: #4a5568;
  }

  .message.error {
    color: #e53e3e;
    font-weight: 500;
  }

  .message.success {
    color: #38a169;
    font-weight: 500;
  }

  .login-link {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    text-decoration: none;
    padding: 0.85rem 2rem;
    border-radius: 12px;
    font-weight: 600;
    font-size: 1rem;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
  }

  .login-link:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(102, 126, 234, 0.5);
  }

  .arrow {
    transition: transform 0.3s ease;
    font-size: 1.2rem;
  }

  .login-link:hover .arrow {
    transform: translateX(5px);
  }

  @media (max-width: 768px) {
    .page-wrapper {
      padding: 1.5rem;
    }

    .card {
      padding: 2rem 1.5rem;
    }

    h1 {
      font-size: 1.5rem;
    }

    .icon {
      font-size: 3rem;
    }

    .spinner {
      width: 50px;
      height: 50px;
    }

    .message {
      font-size: 1rem;
    }
  }
</style>