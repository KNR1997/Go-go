package auth

import (
	"context"
	"errors"

	repo "github.com/knr1997/quiz-tracker-backend/internal/adapters/postgresql/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, params RegisterParams) (repo.User, error)
	Login(ctx context.Context, params LoginParams) (repo.User, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) Register(ctx context.Context, p RegisterParams) (repo.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return repo.User{}, err
	}

	return s.repo.CreateUser(ctx, repo.CreateUserParams{
		Email:        p.Email,
		PasswordHash: string(hash),
	})
}

func (s *svc) Login(ctx context.Context, p LoginParams) (repo.User, error) {
	user, err := s.repo.FindUserByEmail(ctx, p.Email)
	if err != nil {
		return repo.User{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(p.Password),
	); err != nil {
		return repo.User{}, errors.New("invalid credentials")
	}

	return user, nil
}
