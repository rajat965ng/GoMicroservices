package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rajat965ng/p2PostgreSQL/domain"
)

type PostgreSQL struct {
	Pool *pgxpool.Pool
	Ctx  context.Context
}

func NewConection() (*PostgreSQL, error) {
	ctx := context.Background()
	conn := &PostgreSQL{}
	if pool, err := pgxpool.Connect(ctx, "postgres://user:password@localhost:5432/dbname?sslmode=disable"); err != nil {
		fmt.Printf("Error while connect: %s\n", err.Error())
		return nil, err
	} else {
		conn.Pool = pool
		conn.Ctx = ctx
	}
	return conn, nil
}

func (pgSql *PostgreSQL) Close() {
	fmt.Printf("Closing DB connection !!!")
	pgSql.Pool.Close()
}

func (pgSql *PostgreSQL) FindUserById(id string) *domain.Name {
	name := &domain.Name{}
	query := `SELECT nconst, primary_name, birth_year, death_year, primary_professions, known_for_titles FROM "names" WHERE nconst = $1`
	if err := pgSql.Pool.QueryRow(pgSql.Ctx, query,id).Scan(&name.Nconst, &name.Primary_name, &name.Birth_year, &name.Death_year, &name.Primary_professions, &name.Known_for_titles); err != nil {
		fmt.Printf("Error while querying data: %s\n", err.Error())
	}
	return name
}
