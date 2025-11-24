package data

import (
	"database/sql"
	"errors" // You might need to add this import for ErrEditConflict
	"time"

	"github.com/amari03/SysAdmin-FinalProject/Backend/internal/validator"
)

// The MoodNote struct maps directly to the columns in our mood_notes table.
type MoodNote struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Emoji     string    `json:"emoji"`
	Color     string    `json:"color"`
	UserID    int64     `json:"-"` // We use '-' so this field is hidden from JSON responses
	Version   int       `json:"version"` // IMPORTANT: Make sure this is NOT hidden from JSON for updates.
}

// The MoodNoteModel struct will wrap our database connection pool.
type MoodNoteModel struct {
	DB *sql.DB
}

// We also need a validation function for when a user creates a new note.
func ValidateMoodNote(v *validator.Validator, note *MoodNote) {
	v.Check(note.Title != "", "title", "must be provided")
	v.Check(len(note.Title) <= 100, "title", "must not be more than 100 bytes long")

	v.Check(note.Content != "", "content", "must be provided")

	v.Check(note.Emoji != "", "emoji", "must be provided")

	v.Check(note.Color != "", "color", "must be provided")

	v.Check(string(note.Color[0]) == "#", "color", "must start with a #")
	v.Check(len(note.Color) == 7 || len(note.Color) == 4, "color", "must be a valid hex code (e.g., #RRGGBB)")
}

// Insert() adds a new record to the mood_notes table.
func (m MoodNoteModel) Insert(note *MoodNote) error {
	// Define the SQL query for inserting a new record.
	// We use the RETURNING clause to get the system-generated values (id, created_at, version).
	query := `
		INSERT INTO mood_notes (title, content, emoji, color, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`

	// Create an args slice containing the values for the placeholder parameters.
	args := []interface{}{note.Title, note.Content, note.Emoji, note.Color, note.UserID}

	// Use the QueryRow() method to execute the SQL query on the connection pool.
	// This returns the generated values, which we scan into the note struct.
	return m.DB.QueryRow(query, args...).Scan(&note.ID, &note.CreatedAt, &note.Version)
}

// GetAllForUser() returns a slice of all mood notes for a specific user.
func (m MoodNoteModel) GetAllForUser(userID int64) ([]*MoodNote, error) {
	query := `
		SELECT id, created_at, title, content, emoji, color, user_id, version
		FROM mood_notes
		WHERE user_id = $1
		ORDER BY created_at DESC`

	rows, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice to hold the MoodNote structs.
	notes := []*MoodNote{}

	for rows.Next() {
		var note MoodNote
		err := rows.Scan(
			&note.ID,
			&note.CreatedAt,
			&note.Title,
			&note.Content,
			&note.Emoji,
			&note.Color,
			&note.UserID, // Scan the UserID here
			&note.Version,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

// Delete() removes a specific mood note from the database belonging to a specific user.
func (m MoodNoteModel) Delete(id int64, userID int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	// The query now includes a WHERE clause for both the note ID and the user ID.
	// This ensures a user can ONLY delete their own notes.
	query := `
        DELETE FROM mood_notes
        WHERE id = $1 AND user_id = $2`

	// Pass the userID as the second argument to Exec().
	result, err := m.DB.Exec(query, id, userID)
	if err != nil {
		return err
	}

	// If no rows are affected, it means the note either didn't exist or didn't
	// belong to this user. In either case, we return a "not found" error for security.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// ====================================================================================
// START: NEW CODE TO ADD
// ====================================================================================

// Get() retrieves a specific mood note if and only if it belongs to the specified user.
func (m MoodNoteModel) Get(id int64, userID int64) (*MoodNote, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, title, content, emoji, color, user_id, version
		FROM mood_notes
		WHERE id = $1 AND user_id = $2`

	var note MoodNote

	err := m.DB.QueryRow(query, id, userID).Scan(
		&note.ID,
		&note.CreatedAt,
		&note.Title,
		&note.Content,
		&note.Emoji,
		&note.Color,
		&note.UserID,
		&note.Version,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &note, nil
}

// Update() modifies the details of a specific mood note.
// We use the version number for optimistic locking to prevent race conditions.
func (m MoodNoteModel) Update(note *MoodNote) error {
	query := `
		UPDATE mood_notes
		SET title = $1, content = $2, emoji = $3, color = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version`

	args := []interface{}{
		note.Title,
		note.Content,
		note.Emoji,
		note.Color,
		note.ID,
		note.Version, // This is the version number we fetched before updating.
	}

	// Execute the query. If no rows are affected, it means the version number
	// has changed since we last fetched it, indicating an edit conflict.
	err := m.DB.QueryRow(query, args...).Scan(&note.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEditConflict
		}
		return err
	}
	return nil
}

// ====================================================================================
// END: NEW CODE TO ADD
// ====================================================================================