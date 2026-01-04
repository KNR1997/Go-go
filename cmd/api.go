package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
	repo "github.com/knr1997/quiz-tracker-backend/internal/adapters/postgresql/sqlc"
	"github.com/knr1997/quiz-tracker-backend/internal/auth"
	"github.com/knr1997/quiz-tracker-backend/internal/courses"
	"github.com/knr1997/quiz-tracker-backend/internal/orders"
	"github.com/knr1997/quiz-tracker-backend/internal/products"
	"github.com/knr1997/quiz-tracker-backend/internal/quizzes"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// -------- CORS --------
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:1420", // Vue dev
			"http://localhost:3000", // just in case
			"tauri://localhost",     // Tauri desktop
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Link",
		},
		AllowCredentials: true,
		MaxAge:           300, // cache preflight
	}))
	// ----------------------

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes
	r.Use(middleware.Timeout(60 * time.Second))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	courseService := courses.NewService(repo.New(app.db))
	courseHandler := courses.NewHandler(courseService)
	r.Get("/courses", courseHandler.ListCourses)
	r.Get("/courses/{id}", courseHandler.GetCourseByID)
	r.Post("/courses", courseHandler.CreateCourse)
	r.Put("/courses/{id}", courseHandler.UpdateCourse)
	r.Delete("/courses/{id}", courseHandler.DeleteCourse)

	quizService := quizzes.NewService(repo.New(app.db))
	quizHandler := quizzes.NewHandler(quizService)
	r.Get("/quizzes", quizHandler.ListQuizzes)
	r.Get("/quizzes/{id}", quizHandler.GetQuizByID)
	r.Post("/quizzes", quizHandler.CreateQuiz)
	r.Put("/quizzes/{id}", quizHandler.UpdateQuiz)
	r.Delete("/quizzes/{id}", quizHandler.DeleteQuiz)

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	orderHandler := orders.NewHandler(nil)
	r.Post("/orders", orderHandler.PlaceOrder)

	authService := auth.NewService(repo.New(app.db))
	authHandler := auth.NewHandler(authService)
	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	db *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
