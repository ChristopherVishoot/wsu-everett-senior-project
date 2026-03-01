package service

import (
	"database/sql"
	"time"
)

// User represents a user in the database
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Service defines the interface for the service
type Service interface {
	SayHello(name string) string
	CreateUser(name, email string) (*User, error)
}

// service is the implementation of Service
type service struct {
	db *sql.DB
}

// NewService creates a new Service instance
func NewService(db *sql.DB) Service {
	return &service{db: db}
}

// SayHello returns a greeting message
func (s *service) SayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}

// CreateUser inserts a new user into the database
func (s *service) CreateUser(name, email string) (*User, error) {
	var user User
	err := s.db.QueryRow(
		`INSERT INTO users (name, email) VALUES ($1, $2) 
		 RETURNING id, name, email, created_at`,
		name, email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
