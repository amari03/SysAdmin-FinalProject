<script>
  import { onMount } from 'svelte';
  import { authStore, logout } from '../../lib/auth.js';
  import AddNoteModal from '../../lib/AddNoteModal.svelte';
   import ConfirmationModal from '../../lib/ConfirmationModal.svelte';

  let notes = [];
  const user = $authStore.user;
  
  let showModal = false;
  let currentNoteToEdit = null;

  // 2. Add new state variables for the confirmation modal
  let showConfirmModal = false;
  let noteIdToDelete = null; // To remember which note to delete after confirmation

  function showAddNoteModal() {
    currentNoteToEdit = null;
    showModal = true;
  }

  function showEditNoteModal(note) {
    currentNoteToEdit = note;
    showModal = true;
  }

  async function handleSaveNote(event) {
    const noteData = event.detail;
    const token = $authStore.token;

    const payload = {
        title: noteData.title,
        content: noteData.content,
        emoji: noteData.emoji,
        color: noteData.color,
    };

    if (noteData.id) {
        try {
            const response = await fetch(`http://localhost:4000/v1/moodnotes/${noteData.id}`, {
                method: 'PATCH',
                headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
                body: JSON.stringify(payload)
            });
            if (!response.ok) {
                const errorBody = await response.json();
                console.error("Backend validation error:", errorBody);
                throw new Error('Failed to update the note. Check console for details.');
            }
            const updatedNote = (await response.json()).mood_note;
            notes = notes.map(n => n.id === updatedNote.id ? updatedNote : n);
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    } else {
        try {
            const response = await fetch('http://localhost:4000/v1/moodnotes', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
                body: JSON.stringify(payload)
            });
            if (!response.ok) { 
                const errorBody = await response.json();
                console.error("Backend validation error:", errorBody);
                throw new Error('Failed to save the note. Check console for details.');
            }
            const data = await response.json();
            notes = [...notes, data.mood_note];
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    }

    showModal = false;
  }

   // --- Functions for the Delete Confirmation Modal ---

  // 3. This function INITIATES the delete process
  function promptToDelete(noteId) {
    noteIdToDelete = noteId; // Remember the ID
    showConfirmModal = true; // Show our new modal
  }

  // 4. This function EXECUTES the delete after the user confirms in the modal
  async function confirmDelete() {
    if (!noteIdToDelete) return; // Safety check

    const token = $authStore.token;
    try {
      const response = await fetch(`http://localhost:4000/v1/moodnotes/${noteIdToDelete}`, {
        method: 'DELETE',
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (!response.ok) { throw new Error('Failed to delete note.'); }
      // Update the UI
      notes = notes.filter(note => note.id !== noteIdToDelete);
    } catch (error) {
      console.error(error);
      alert(error.message);
    } finally {
      // 5. Reset the state and hide the modal
      showConfirmModal = false;
      noteIdToDelete = null;
    }
  }


  onMount(async () => {
    const token = $authStore.token;
    try {
      const response = await fetch('http://localhost:4000/v1/moodnotes', {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (!response.ok) { throw new Error('Failed to fetch notes.'); }
      const data = await response.json();
      notes = data.mood_notes;
    } catch (error) {
      console.error(error);
      alert(error.message);
    }
  });
</script>

<main>
  <header>
    <div class="header-content">
      <div class="welcome-section">
        <h1>Hello, {user?.name}! 👋</h1>
        <p class="subtitle">How are you feeling today?</p>
      </div>
      <button class="logout-btn" on:click={logout}>
        <span>Log Out</span>
      </button>
    </div>
  </header>

  <div class="notes-grid">
    <button class="add-note-tile" on:click={showAddNoteModal}>
      <div class="add-content">
        <span class="plus-icon">+</span>
        <span class="add-text">New Note</span>
      </div>
    </button>

    {#each notes as note (note.id)}
      <div class="note-card" style="background-color: {note.color};">
        <div class="note-header">
          <div class="note-emoji">{note.emoji}</div>
          <h2>{note.title}</h2>
        </div>
        <p class="note-content">{note.content}</p>
        <div class="note-footer">
          <small>{new Date(note.created_at).toLocaleDateString()}</small>
          <div class="note-actions">
            <button class="icon-btn edit-btn" on:click={() => showEditNoteModal(note)} title="Edit">
              ✏️
            </button>
            <button class="icon-btn delete-btn" on:click={() => promptToDelete(note.id)} title="Delete">
              🗑️
            </button>
          </div>
        </div>
      </div>
    {/each}
  </div>

  {#if showModal}
    <AddNoteModal 
      noteToEdit={currentNoteToEdit}
      on:save={handleSaveNote}
      on:cancel={() => showModal = false}
    />
  {/if}

   <!-- 7. Add our new Confirmation Modal here -->
  {#if showConfirmModal}
    <ConfirmationModal
      title="Delete Note"
      message="Are you sure you want to permanently delete this note?"
      on:confirm={confirmDelete}
      on:cancel={() => showConfirmModal = false}
    />
  {/if}

</main>

<style>
  :global(body) {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    margin: 0;
    padding: 0;
  }

  main {
    padding: 2rem;
    max-width: 1400px;
    margin: 0 auto;
  }

  header {
    margin-bottom: 3rem;
  }

  .header-content {
    background: white;
    border-radius: 20px;
    padding: 2rem 2.5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  }

  .welcome-section h1 {
    margin: 0 0 0.5rem 0;
    font-size: 2rem;
    font-weight: 700;
    color: #1a202c;
  }

  .subtitle {
    margin: 0;
    color: #718096;
    font-size: 1rem;
  }

  .logout-btn {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 12px;
    padding: 0.75rem 1.5rem;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
  }

  .logout-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(102, 126, 234, 0.5);
  }

  .notes-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 1.5rem;
  }

  .add-note-tile {
    background: rgba(255, 255, 255, 0.95);
    border: 3px dashed #cbd5e0;
    border-radius: 20px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 250px;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  }

  .add-note-tile:hover {
    background: white;
    border-color: #667eea;
    transform: translateY(-5px);
    box-shadow: 0 8px 25px rgba(102, 126, 234, 0.3);
  }

  .add-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.5rem;
  }

  .plus-icon {
    font-size: 4rem;
    color: #cbd5e0;
    transition: color 0.3s ease;
  }

  .add-note-tile:hover .plus-icon {
    color: #667eea;
  }

  .add-text {
    font-size: 1.1rem;
    font-weight: 600;
    color: #a0aec0;
    transition: color 0.3s ease;
  }

  .add-note-tile:hover .add-text {
    color: #667eea;
  }

  .note-card {
    border-radius: 20px;
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    min-height: 250px;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.15);
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
  }

  .note-card::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: linear-gradient(135deg, rgba(255,255,255,0.2) 0%, rgba(255,255,255,0) 100%);
    pointer-events: none;
  }

  .note-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2);
  }

  .note-header {
    display: flex;
    align-items: center;
    margin-bottom: 1rem;
    gap: 0.75rem;
  }

  .note-emoji {
    font-size: 2rem;
    line-height: 1;
  }

  .note-header h2 {
    margin: 0;
    font-size: 1.3rem;
    font-weight: 700;
    color: #1a202c;
    flex-grow: 1;
  }

  .note-content {
    flex-grow: 1;
    font-size: 1rem;
    line-height: 1.6;
    color: #2d3748;
    word-wrap: break-word;
    margin: 0 0 1rem 0;
  }

  .note-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-top: 1rem;
    border-top: 1px solid rgba(0, 0, 0, 0.1);
  }

  .note-footer small {
    font-size: 0.85rem;
    color: #4a5568;
    font-weight: 500;
  }

  .note-actions {
    display: flex;
    gap: 0.5rem;
  }

  .icon-btn {
    background: rgba(255, 255, 255, 0.8);
    border: none;
    border-radius: 8px;
    cursor: pointer;
    font-size: 1.1rem;
    padding: 0.4rem 0.6rem;
    transition: all 0.2s ease;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .icon-btn:hover {
    background: white;
    transform: scale(1.1);
  }

  .edit-btn:hover {
    box-shadow: 0 2px 8px rgba(102, 126, 234, 0.4);
  }

  .delete-btn:hover {
    box-shadow: 0 2px 8px rgba(239, 68, 68, 0.4);
  }

  @media (max-width: 768px) {
    main {
      padding: 1.5rem;
    }

    .header-content {
      flex-direction: column;
      gap: 1.5rem;
      text-align: center;
    }

    .welcome-section h1 {
      font-size: 1.5rem;
    }

    .notes-grid {
      grid-template-columns: 1fr;
    }
  }
</style>