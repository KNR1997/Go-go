package auth

import (
	"net/http"

	"github.com/knr1997/quiz-tracker-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

// POST /auth/register
func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var params RegisterParams
	if err := json.Read(r, &params); err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if params.Email == "" || params.Password == "" {
		json.Write(w, http.StatusBadRequest, map[string]string{"error": "email and password required"})
		return
	}

	user, err := h.service.Register(r.Context(), params)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	token, _ := GenerateToken(user.ID)

	json.Write(w, http.StatusCreated, map[string]string{
		"token": token,
	})
}

// POST /auth/login
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
	if err := json.Read(r, &params); err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	user, err := h.service.Login(r.Context(), params)
	if err != nil {
		json.Write(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		return
	}

	token, _ := GenerateToken(user.ID)

	json.Write(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
