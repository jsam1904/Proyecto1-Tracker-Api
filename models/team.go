package models

import "time"

type Team struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Stadium   string    `json:"stadium"`
	Founded   int       `json:"founded"`
	ImageURL  *string   `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTeamRequest struct {
	Name     string  `json:"name"`
	City     string  `json:"city"`
	Stadium  string  `json:"stadium"`
	Founded  int     `json:"founded"`
	ImageURL *string `json:"image_url"`
}

type UpdateTeamRequest struct {
	Name     *string `json:"name"`
	City     *string `json:"city"`
	Stadium  *string `json:"stadium"`
	Founded  *int    `json:"founded"`
	ImageURL *string `json:"image_url"`
}

type PaginatedResponse struct {
	Data       []Team `json:"data"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
}

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}
