package quizzes

import (
	"context"

	repo "github.com/knr1997/quiz-tracker-backend/internal/adapters/postgresql/sqlc"
)

type createQuizParams struct {
	CourseID   int64  `json:"courseId"`
	WeekNumber int    `json:"weekNumber"`
	DateTime   string `json:"dataTime"`
	Status     string `json:"status"`
}

type updateQuizParams struct {
	WeekNumber int    `json:"weekNumber"`
	DateTime   string `json:"dataTime"`
	Status     string `json:"status"`
}

type QuizWithCourse struct {
}

type Service interface {
	ListQuizzes(ctx context.Context) ([]repo.Quiz, error)
	ListQuizzesWithCourse(ctx context.Context) ([]repo.ListQuizzesWithCourseRow, error)
	GetQuizByID(ctx context.Context, id int64) (repo.Quiz, error)
	CreateQuiz(ctx context.Context, tempQuiz createQuizParams) (repo.Quiz, error)
	UpdateQuiz(ctx context.Context, id int64, tempQuiz updateQuizParams) (repo.Quiz, error)
	DeleteQuiz(ctx context.Context, id int64) error
}
