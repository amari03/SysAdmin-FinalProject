package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/amari03/SysAdmin-FinalProject/Backend/internal/data"
	"github.com/amari03/SysAdmin-FinalProject/Backend/internal/validator"
)

// The createMoodNoteHandler function will handle POST requests to /v1/moodnotes.
func (a *applicationDependencies) createMoodNoteHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Define a struct to hold the input from the frontend.
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Emoji   string `json:"emoji"`
		Color   string `json:"color"`
	}

	// 2. Read the JSON from the request body into the input struct.
	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	// 3. Get the currently logged-in user from the request context.
	// This is the most important step for security!
	user := a.contextGetUser(r)

	// 4. Copy the data from the input struct into a new MoodNote model.
	note := &data.MoodNote{
		Title:   input.Title,
		Content: input.Content,
		Emoji:   input.Emoji,
		Color:   input.Color,
		UserID:  user.ID, // Associate the note with the logged-in user.
	}

	// 5. Validate the new note.
	v := validator.New()
	if data.ValidateMoodNote(v, note); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	// 6. Insert the new note into the database.
	err = a.models.MoodNotes.Insert(note)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// 7. Send a 201 Created response back to the user with the new note's data.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/moodnotes/%d", note.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"mood_note": note}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

// listMoodNotesHandler retrieves all mood notes for the authenticated user.
func (a *applicationDependencies) listMoodNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the current user.
	user := a.contextGetUser(r)

	// Call the model method to get all notes for this user.
	notes, err := a.models.MoodNotes.GetAllForUser(user.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// Send the slice of notes as a JSON response.
	err = a.writeJSON(w, http.StatusOK, envelope{"mood_notes": notes}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

// deleteMoodNoteHandler deletes a specific mood note.
func (a *applicationDependencies) deleteMoodNoteHandler(w http.ResponseWriter, r *http.Request) {
    // Get the note ID from the URL.
    id, err := a.readIDParam(r)
    if err != nil {
        a.notFoundResponse(w, r)
        return
    }

    // Get the current user from the request context. This is the crucial security step.
    user := a.contextGetUser(r)

    // Call the model to delete the note, passing both the note ID and the user's ID.
    err = a.models.MoodNotes.Delete(id, user.ID)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            // We return a 404 Not Found if the note doesn't exist OR the user doesn't own it.
            a.notFoundResponse(w, r)
        default:
            a.serverErrorResponse(w, r, err)
        }
        return
    }

    // Send a 200 OK response with a success message.
    err = a.writeJSON(w, http.StatusOK, envelope{"message": "mood note successfully deleted"}, nil)
    if err != nil {
        a.serverErrorResponse(w, r, err)
    }
}

// updateMoodNoteHandler handles PATCH requests to update a mood note.
func (a *applicationDependencies) updateMoodNoteHandler(w http.ResponseWriter, r *http.Request) {
    // Get the note ID from the URL.
    id, err := a.readIDParam(r)
    if err != nil {
        a.notFoundResponse(w, r)
        return
    }

    // Get the current user.
    user := a.contextGetUser(r)

    // Fetch the existing note from the database to ensure it exists and the user owns it.
    note, err := a.models.MoodNotes.Get(id, user.ID)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            a.notFoundResponse(w, r)
        default:
            a.serverErrorResponse(w, r, err)
        }
        return
    }

    // Define a struct to hold the input from the frontend.
    // We use pointers to distinguish between a field being absent vs. having a zero value.
    var input struct {
        Title   *string `json:"title"`
        Content *string `json:"content"`
        Emoji   *string `json:"emoji"`
        Color   *string `json:"color"`
    }

    // Read the JSON from the request body.
    err = a.readJSON(w, r, &input)
    if err != nil {
        a.badRequestResponse(w, r, err)
        return
    }
    
    // Update the note's fields if the corresponding key exists in the JSON.
    if input.Title != nil {
        note.Title = *input.Title
    }
    if input.Content != nil {
        note.Content = *input.Content
    }
    if input.Emoji != nil {
        note.Emoji = *input.Emoji
    }
    if input.Color != nil {
        note.Color = *input.Color
    }
    
    // Validate the updated note.
    v := validator.New()
    if data.ValidateMoodNote(v, note); !v.IsEmpty() {
        a.failedValidationResponse(w, r, v.Errors)
        return
    }

    // Save the updated note back to the database.
    err = a.models.MoodNotes.Update(note)
    if err != nil {
        a.serverErrorResponse(w, r, err)
        return
    }

    // Send the updated note back to the client.
    err = a.writeJSON(w, http.StatusOK, envelope{"mood_note": note}, nil)
    if err != nil {
        a.serverErrorResponse(w, r, err)
    }
}