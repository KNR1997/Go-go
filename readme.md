# Quiz Tracker Backend

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

A robust Golang backend API for managing and tracking quiz applications. Provides RESTful endpoints for quiz management, user tracking, and analytics.

## Clone the Repository

```
$ git clone https://github.com/KNR1997/quiz-tracker-backend.git
$ cd quiz-tracker-backend
```

## [Start the Database with Docker](start-the-database-with-docker/)

```
$ docker compose up
```

## [Run Database Migrations](run-database-migrations/)

```
$ goose up
```

## [Generate SQLC Code](generate-sqlc-code/)

```
$ goose -s create create_users sql
```

## [Generate SQLC Code](generate-sqlc-code/)

```
$ sqlc generate
```

## [Start the Server](start-the-server/)

```
$ go run ./cmd
```
