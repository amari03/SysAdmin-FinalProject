<script>
  import { createEventDispatcher } from 'svelte';

  // Props to make the modal reusable
  export let title = 'Confirm Action';
  export let message = 'Are you sure you want to proceed?';

  const dispatch = createEventDispatcher();

  // These functions will notify the parent component of the user's choice.
  function handleConfirm() {
    dispatch('confirm');
  }

  function handleCancel() {
    dispatch('cancel');
  }
</script>

<div class="modal-backdrop" on:click={handleCancel}>
  <div class="modal-content" on:click|stopPropagation>
    <div class="modal-header">
      <div class="icon-wrapper">
        <span class="warning-icon">⚠️</span>
      </div>
      <h2>{title}</h2>
    </div>
    
    <div class="modal-body">
      <p>{message}</p>
    </div>
    
    <div class="button-group">
      <button class="cancel-btn" on:click={handleCancel}>Cancel</button>
      <button class="confirm-btn" on:click={handleConfirm}>Confirm</button>
    </div>
  </div>
</div>

<style>
  .modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(4px);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 100;
    animation: fadeIn 0.2s ease;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  .modal-content {
    background-color: white;
    padding: 0;
    border-radius: 20px;
    width: 90%;
    max-width: 450px;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
    animation: slideUp 0.3s ease;
    overflow: hidden;
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

  .modal-header {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    padding: 2rem;
    text-align: center;
  }

  .icon-wrapper {
    margin-bottom: 1rem;
  }

  .warning-icon {
    font-size: 3.5rem;
    display: inline-block;
    animation: pulse 2s ease-in-out infinite;
  }

  @keyframes pulse {
    0%, 100% {
      transform: scale(1);
    }
    50% {
      transform: scale(1.1);
    }
  }

  h2 {
    margin: 0;
    font-size: 1.5rem;
    font-weight: 700;
    color: white;
  }

  .modal-body {
    padding: 2rem;
  }

  p {
    margin: 0;
    font-size: 1.05rem;
    line-height: 1.6;
    color: #4a5568;
    text-align: center;
  }

  .button-group {
    display: flex;
    justify-content: center;
    gap: 0.75rem;
    padding: 0 2rem 2rem;
  }
  
  .cancel-btn, .confirm-btn {
    padding: 0.85rem 2rem;
    border-radius: 12px;
    border: none;
    cursor: pointer;
    font-weight: 600;
    font-size: 1rem;
    transition: all 0.3s ease;
    flex: 1;
    max-width: 150px;
  }
  
  .cancel-btn {
    background-color: #edf2f7;
    color: #4a5568;
  }

  .cancel-btn:hover {
    background-color: #e2e8f0;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }
  
  .confirm-btn {
    background: linear-gradient(135deg, #fc8181 0%, #e53e3e 100%);
    color: white;
    box-shadow: 0 4px 15px rgba(229, 62, 62, 0.4);
  }

  .confirm-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(229, 62, 62, 0.5);
  }

  @media (max-width: 768px) {
    .modal-content {
      width: 95%;
    }

    .modal-header {
      padding: 1.5rem;
    }

    .warning-icon {
      font-size: 3rem;
    }

    h2 {
      font-size: 1.25rem;
    }

    .modal-body {
      padding: 1.5rem;
    }

    p {
      font-size: 1rem;
    }

    .button-group {
      padding: 0 1.5rem 1.5rem;
      gap: 0.5rem;
    }

    .cancel-btn, .confirm-btn {
      padding: 0.75rem 1.5rem;
      font-size: 0.95rem;
    }
  }
</style>