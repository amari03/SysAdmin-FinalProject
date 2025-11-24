<script>
  import { onMount, createEventDispatcher } from 'svelte';

  export let noteToEdit = null;

  // Internal state of the form
  let id = null;
  let version = null;
  let title = '';
  let content = '';
  let emoji = '😊';
  let color = '#FFF9A5'; // Default color

  const emotions = [
    { emoji: '😊', name: 'Happy' },
    { emoji: '😢', name: 'Sad' },
    { emoji: '😠', name: 'Angry' },
    { emoji: '😮', name: 'Surprised' },
    { emoji: '🤔', name: 'Thinking' },
    { emoji: '😎', name: 'Cool' },
  ];

  onMount(() => {
    // If a note was passed in, we are in "edit mode".
    if (noteToEdit) {
      id = noteToEdit.id;
      version = noteToEdit.version;
      title = noteToEdit.title;
      content = noteToEdit.content;
      emoji = noteToEdit.emoji;
      color = noteToEdit.color;
    }
  });

  const dispatch = createEventDispatcher();

  function handleSave() {
    dispatch('save', { id, version, title, content, emoji, color });
  }

  function handleCancel() {
    dispatch('cancel');
  }
</script>

<div class="modal-backdrop" on:click={handleCancel}>
  <div class="modal-content" on:click|stopPropagation>
    <div class="modal-header">
      <h2>{noteToEdit ? 'Edit Mood Note' : 'Add a New Mood Note'}</h2>
      <button class="close-btn" on:click={handleCancel}>✕</button>
    </div>

    <div class="form-group">
      <label for="title">Title</label>
      <input type="text" id="title" bind:value={title} placeholder="A title for your note...">
    </div>

    <div class="form-group">
      <label for="content">Description</label>
      <textarea id="content" rows="4" bind:value={content} placeholder="How are you feeling today?"></textarea>
    </div>

    <div class="form-group">
      <label>How are you feeling?</label>
      <div class="emoji-selector">
        {#each emotions as emotion (emotion.emoji)}
          <button
            type="button"
            class="emoji-btn"
            class:selected={emoji === emotion.emoji}
            on:click={() => (emoji = emotion.emoji)}
            title={emotion.name}
          >
            {emotion.emoji}
          </button>
        {/each}
      </div>
    </div>

    <div class="form-group">
      <label for="color">Choose a color</label>
      <div class="color-picker-wrapper">
        <input type="color" id="color" bind:value={color}>
        <span class="color-preview" style="background-color: {color};"></span>
      </div>
    </div>

    <div class="button-group">
      <button class="cancel-btn" on:click={handleCancel}>Cancel</button>
      <button class="save-btn" on:click={handleSave}>
        {noteToEdit ? 'Update Note' : 'Save Note'}
      </button>
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
    max-width: 540px;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
    animation: slideUp 0.3s ease;
    max-height: 90vh;
    overflow-y: auto;
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
    padding: 1.75rem 2rem;
    border-radius: 20px 20px 0 0;
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
  }

  .modal-header h2 {
    margin: 0;
    font-size: 1.5rem;
    font-weight: 700;
    color: white;
  }

  .close-btn {
    background: rgba(255, 255, 255, 0.2);
    border: none;
    color: white;
    font-size: 1.5rem;
    width: 36px;
    height: 36px;
    border-radius: 50%;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
    padding: 0;
    line-height: 1;
  }

  .close-btn:hover {
    background: rgba(255, 255, 255, 0.3);
    transform: rotate(90deg);
  }

  .form-group {
    margin-bottom: 1.5rem;
    padding: 0 2rem;
  }

  label {
    display: block;
    margin-bottom: 0.6rem;
    font-weight: 600;
    color: #2d3748;
    font-size: 0.95rem;
  }

  input[type="text"], textarea {
    width: 100%;
    padding: 0.85rem;
    border: 2px solid #e2e8f0;
    border-radius: 12px;
    box-sizing: border-box;
    font-family: inherit;
    font-size: 1rem;
    transition: all 0.2s ease;
    background: #f7fafc;
  }

  input[type="text"]:focus, textarea:focus {
    outline: none;
    border-color: #667eea;
    background: white;
    box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
  }

  textarea {
    resize: vertical;
    min-height: 100px;
  }

  .emoji-selector {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
  }

  .emoji-btn {
    font-size: 2rem;
    background: #f7fafc;
    border: 3px solid transparent;
    border-radius: 15px;
    padding: 0.75rem;
    cursor: pointer;
    transition: all 0.2s ease;
    width: 60px;
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .emoji-btn:hover {
    background-color: #edf2f7;
    transform: scale(1.05);
  }

  .emoji-btn.selected {
    border-color: #667eea;
    background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.2);
  }

  .color-picker-wrapper {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  input[type="color"] {
    width: 80px;
    height: 50px;
    border: 3px solid #e2e8f0;
    border-radius: 12px;
    cursor: pointer;
    padding: 4px;
    background: white;
    transition: all 0.2s ease;
  }

  input[type="color"]:hover {
    border-color: #667eea;
  }

  .color-preview {
    width: 50px;
    height: 50px;
    border-radius: 12px;
    border: 3px solid #e2e8f0;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }

  .button-group {
    display: flex;
    justify-content: flex-end;
    gap: 0.75rem;
    padding: 1.5rem 2rem 2rem;
  }

  .save-btn, .cancel-btn {
    padding: 0.85rem 1.75rem;
    border-radius: 12px;
    border: none;
    cursor: pointer;
    font-weight: 600;
    font-size: 1rem;
    transition: all 0.3s ease;
  }

  .save-btn {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
  }

  .save-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(102, 126, 234, 0.5);
  }

  .cancel-btn {
    background-color: #edf2f7;
    color: #4a5568;
  }

  .cancel-btn:hover {
    background-color: #e2e8f0;
  }

  @media (max-width: 768px) {
    .modal-content {
      width: 95%;
      max-height: 95vh;
    }

    .modal-header {
      padding: 1.5rem;
    }

    .modal-header h2 {
      font-size: 1.25rem;
    }

    .form-group {
      padding: 0 1.5rem;
    }

    .button-group {
      padding: 1.5rem;
      flex-direction: column-reverse;
    }

    .save-btn, .cancel-btn {
      width: 100%;
    }

    .emoji-btn {
      width: 55px;
      height: 55px;
      font-size: 1.75rem;
    }
  }
</style>