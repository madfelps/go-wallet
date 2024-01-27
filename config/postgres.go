package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitializePostgres() (*sql.DB, error) {
	logger := GetLogger("postgres")
	// Check if the database file exists

	// Create DB and connect
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}

	/* sql := `CREATE TABLE wallets (
		name VARCHAR(255),
		balance VARCHAR(255)
	);`
	_, errDb := db.Exec(sql)
	if errDb != nil {
		fmt.Println("Table init error")
	}
	fmt.Println("Table created") */

	return db, nil
}
