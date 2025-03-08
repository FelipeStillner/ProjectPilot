package adapter

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	c "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"
	_ "github.com/lib/pq"
)

type TeamRepository struct {
	db *sql.DB
}

func NewTeamRepository() *TeamRepository {
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

	return &TeamRepository{db: db}
}

func (t *TeamRepository) Create(team c.Team) (*c.Team, error) {
	stmt, err := t.db.Prepare("INSERT INTO team (id, name, created_at, updated_at) VALUES ($1, $2, $3, $3)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(team.Id, team.Name, time.Now())
	if err != nil {
		return nil, err
	}

	return &team, nil
}
