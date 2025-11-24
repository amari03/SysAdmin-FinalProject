package data

import "database/sql"

type Models struct {
    Users UserModel
    Tokens TokenModel
    MoodNotes MoodNoteModel
}

func NewModels(db *sql.DB) Models {
    return Models{
        Users: UserModel{DB: db},
        Tokens: TokenModel{DB: db},
        MoodNotes: MoodNoteModel{DB: db},
    }
}