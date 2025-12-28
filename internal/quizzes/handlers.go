package quizzes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/knr1997/quiz-tracker-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzes, err := h.service.ListQuizzesWithCourse(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, quizzes)
}

func (h *handler) GetQuizByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid quiz ID",
		})
		return
	}

	quiz, err := h.service.GetQuizByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		// You can add more specific error handling here
		status := http.StatusInternalServerError
		errorMsg := err.Error()

		// Example: Check for "not found" errors
		// if errors.Is(err, sql.ErrNoRows) {
		//     status = http.StatusNotFound
		//     errorMsg = "Course not found"
		// }

		json.Write(w, status, map[string]string{
			"error": errorMsg,
		})
		return
	}

	json.Write(w, http.StatusOK, quiz)
}

func (h *handler) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var tempQuiz createQuizParams
	if err := json.Read(r, &tempQuiz); err != nil {
		log.Println(err)
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	// Validate input
	// if tempQuiz.Name == "" || tempQuiz.Code == "" {
	// 	json.Write(w, http.StatusBadRequest, map[string]string{
	// 		"error": "Name and Code are required",
	// 	})
	// 	return
	// }

	createdCourse, err := h.service.CreateQuiz(r.Context(), tempQuiz)
	if err != nil {
		log.Println(err)
		json.Write(w, http.StatusInternalServerError, map[string]string{
			"error": "Failed to create course",
		})
		return
	}

	json.Write(w, http.StatusCreated, createdCourse)
}

func (h *handler) UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid course ID",
		})
		return
	}

	var tempQuiz updateQuizParams
	if err := json.Read(r, &tempQuiz); err != nil {
		log.Println(err)
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	// if tempQuiz.Name == "" && tempQuiz.Code == "" {
	// 	json.Write(w, http.StatusBadRequest, map[string]string{
	// 		"error": "At least one field (Name or Code) must be provided",
	// 	})
	// 	return
	// }

	updatedCourse, err := h.service.UpdateQuiz(r.Context(), id, tempQuiz)
	if err != nil {
		log.Println(err)
		status := http.StatusInternalServerError
		errorMsg := "Failed to update course"

		// You can add specific error handling here
		// if errors.Is(err, sql.ErrNoRows) {
		//     status = http.StatusNotFound
		//     errorMsg = "Course not found"
		// }

		json.Write(w, status, map[string]string{
			"error": errorMsg,
		})
		return
	}

	json.Write(w, http.StatusOK, updatedCourse)
}

func (h *handler) DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid course ID",
		})
		return
	}

	err = h.service.DeleteQuiz(r.Context(), id)
	if err != nil {
		log.Println(err)
		status := http.StatusInternalServerError
		errorMsg := "Failed to delete course"

		// You can add specific error handling here
		// if errors.Is(err, sql.ErrNoRows) {
		//     status = http.StatusNotFound
		//     errorMsg = "Course not found"
		// }

		json.Write(w, status, map[string]string{
			"error": errorMsg,
		})
		return
	}

	// Return 204 No Content for successful deletion
	w.WriteHeader(http.StatusNoContent)
}

// Helper function to extract ID from request
func getIDFromRequest(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.ParseInt(idStr, 10, 64)
}
