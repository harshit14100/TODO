package models

import "time"

// this is the structure of a single todo item in this app
type Todo struct {
	Id           string `json:"id" db:"id"`
	UserId       string `json:"user_id" db:"user_id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	IsCompleted  bool   `json:"is_completed0" db:"is_completed"`
	IsIncomplete bool   `json:"is_incomplete" db:"is_incomplete"`
	IsPending    bool   `json:"is_pending" db:"is_pending"`
	ExpiresAt    string `json:"expires_at" db:"expires_at"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

// client side data that is recieved provided by client
type CreateTodo struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ExpiresAt   string `json:"expires_at" binding:"required"`
}

type UpdateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted *bool  `json:"is_completed"`
}

func (t *Todo) SyncStatus() {
	if t.IsCompleted {
		t.IsIncomplete = false
		t.IsPending = false
		return
	}

	expiry, err := time.Parse(time.RFC3339, t.ExpiresAt)
	if err != nil {
		t.IsIncomplete = true
		t.IsPending = false
		return
	}

	if time.Now().After(expiry) {
		t.IsPending = true
		t.IsIncomplete = false
	} else {
		t.IsPending = false
		t.IsIncomplete = true
	}
}
