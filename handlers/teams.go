package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jsam1904/Proyecto1-Tracker-Api/db"
	"github.com/jsam1904/Proyecto1-Tracker-Api/models"

	"github.com/go-chi/chi/v5"
)

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string, details map[string]string) {
	respondJSON(w, status, models.ErrorResponse{Error: message, Details: details})
}

// GET /teams
func GetTeams(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	sortField := r.URL.Query().Get("sort")
	sortOrder := strings.ToUpper(r.URL.Query().Get("order"))
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, limit := 1, 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 1000 {
		limit = l
	}

	allowedSorts := map[string]bool{"name": true, "city": true, "founded": true, "stadium": true, "id": true}
	if sortField != "" && !allowedSorts[sortField] {
		respondError(w, http.StatusBadRequest, "Invalid sort field. Allowed: name, city, founded, stadium", nil)
		return
	}
	if sortOrder != "" && sortOrder != "ASC" && sortOrder != "DESC" {
		respondError(w, http.StatusBadRequest, "Invalid order. Use asc or desc", nil)
		return
	}
	if sortField == "" {
		sortField = "id"
	}
	if sortOrder == "" {
		sortOrder = "ASC"
	}

	offset := (page - 1) * limit

	countQuery := "SELECT COUNT(*) FROM teams"
	var countArgs []any
	if q != "" {
		countQuery += " WHERE name ILIKE $1"
		countArgs = append(countArgs, "%"+q+"%")
	}
	var total int
	db.DB.QueryRow(countQuery, countArgs...).Scan(&total)

	query := "SELECT id, name, city, stadium, founded, image_url, created_at, updated_at FROM teams"
	var args []any
	argIdx := 1

	if q != "" {
		query += " WHERE name ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+q+"%")
		argIdx++
	}
	query += " ORDER BY " + sortField + " " + sortOrder
	query += " LIMIT $" + strconv.Itoa(argIdx) + " OFFSET $" + strconv.Itoa(argIdx+1)
	args = append(args, limit, offset)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Database error", nil)
		return
	}
	defer rows.Close()

	teams := []models.Team{}
	for rows.Next() {
		var t models.Team
		rows.Scan(&t.ID, &t.Name, &t.City, &t.Stadium, &t.Founded, &t.ImageURL, &t.CreatedAt, &t.UpdatedAt)
		teams = append(teams, t)
	}

	totalPages := (total + limit - 1) / limit
	respondJSON(w, http.StatusOK, models.PaginatedResponse{
		Data: teams, Page: page, Limit: limit, Total: total, TotalPages: totalPages,
	})
}

// GET /teams/:id
func GetTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var t models.Team
	err := db.DB.QueryRow(
		"SELECT id, name, city, stadium, founded, image_url, created_at, updated_at FROM teams WHERE id = $1", id,
	).Scan(&t.ID, &t.Name, &t.City, &t.Stadium, &t.Founded, &t.ImageURL, &t.CreatedAt, &t.UpdatedAt)

	if err == sql.ErrNoRows {
		respondError(w, http.StatusNotFound, "Team not found", nil)
		return
	}
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Database error", nil)
		return
	}
	respondJSON(w, http.StatusOK, t)
}

// POST /teams
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	details := map[string]string{}
	if strings.TrimSpace(req.Name) == "" {
		details["name"] = "Name is required"
	}
	if strings.TrimSpace(req.City) == "" {
		details["city"] = "City is required"
	}
	if strings.TrimSpace(req.Stadium) == "" {
		details["stadium"] = "Stadium is required"
	}
	if req.Founded < 1800 || req.Founded > time.Now().Year() {
		details["founded"] = "Founded must be a valid year between 1800 and today"
	}
	if len(details) > 0 {
		respondError(w, http.StatusBadRequest, "Validation failed", details)
		return
	}

	var t models.Team
	err := db.DB.QueryRow(
		`INSERT INTO teams (name, city, stadium, founded, image_url)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, name, city, stadium, founded, image_url, created_at, updated_at`,
		req.Name, req.City, req.Stadium, req.Founded, req.ImageURL,
	).Scan(&t.ID, &t.Name, &t.City, &t.Stadium, &t.Founded, &t.ImageURL, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not create team", nil)
		return
	}
	respondJSON(w, http.StatusCreated, t)
}

// PUT /teams/:id
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var exists bool
	db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM teams WHERE id = $1)", id).Scan(&exists)
	if !exists {
		respondError(w, http.StatusNotFound, "Team not found", nil)
		return
	}

	var req models.UpdateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	details := map[string]string{}
	if req.Name != nil && strings.TrimSpace(*req.Name) == "" {
		details["name"] = "Name cannot be empty"
	}
	if req.City != nil && strings.TrimSpace(*req.City) == "" {
		details["city"] = "City cannot be empty"
	}
	if req.Stadium != nil && strings.TrimSpace(*req.Stadium) == "" {
		details["stadium"] = "Stadium cannot be empty"
	}
	if req.Founded != nil && (*req.Founded < 1800 || *req.Founded > time.Now().Year()) {
		details["founded"] = "Founded must be a valid year between 1800 and today"
	}
	if len(details) > 0 {
		respondError(w, http.StatusBadRequest, "Validation failed", details)
		return
	}

	var t models.Team
	err := db.DB.QueryRow(
		`UPDATE teams SET
			name      = COALESCE($1, name),
			city      = COALESCE($2, city),
			stadium   = COALESCE($3, stadium),
			founded   = COALESCE($4, founded),
			image_url = COALESCE($5, image_url),
			updated_at = NOW()
		 WHERE id = $6
		 RETURNING id, name, city, stadium, founded, image_url, created_at, updated_at`,
		req.Name, req.City, req.Stadium, req.Founded, req.ImageURL, id,
	).Scan(&t.ID, &t.Name, &t.City, &t.Stadium, &t.Founded, &t.ImageURL, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not update team", nil)
		return
	}
	respondJSON(w, http.StatusOK, t)
}

// DELETE /teams/:id
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := db.DB.Exec("DELETE FROM teams WHERE id = $1", id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete team", nil)
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		respondError(w, http.StatusNotFound, "Team not found", nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
