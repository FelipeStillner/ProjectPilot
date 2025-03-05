package adapter

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	c "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	dbHost := os.Getenv("HOST_POSTGRES")
	dbPort := os.Getenv("PORT_POSTGRES")
	dbUser := os.Getenv("USER_POSTGRES")
	dbPassword := os.Getenv("PASSWORD_POSTGRES")
	dbName := os.Getenv("DATABASE_POSTGRES")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &UserRepository{db: db}
}

func (t *UserRepository) Create(user c.User) (*c.User, error) {
	stmt, err := t.db.Prepare("INSERT INTO \"user\" (id, username, password, team_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.Id, user.Username, user.Password, user.TeamId, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (t *UserRepository) Read(username string) (*c.User, error) {
	query := "SELECT id, username, password, team_id FROM \"user\" WHERE username = $1"
	row := t.db.QueryRow(query, username)

	var user c.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.TeamId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
