package models

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
type createTodo struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ExpiresAt   string `json:"expires_at" binding:"required"`
}
