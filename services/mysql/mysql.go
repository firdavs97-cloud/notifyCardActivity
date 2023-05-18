package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Iface interface {
	Insert(string, ...interface{}) (int64, error)
	Update(string, ...interface{}) error
	QueryRow(string, int64, ...interface{}) error
	QueryIds(string, ...interface{}) ([]int64, error)
}

// DB represents the MySQL database connection
type DB struct {
	Conn *sql.DB
}

func (db DB) Insert(s string, i ...interface{}) (int64, error) {
	stmt, err := db.Conn.Prepare(s)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	result, err := stmt.Exec(i...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute SQL statement: %v", err)
	}

	// Retrieve the inserted user's primary key (ID)
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	return id, nil
}
func (db DB) Update(s string, i ...interface{}) error {
	stmt, err := db.Conn.Prepare(s)
	if err != nil {
		return fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(i...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL statement: %v", err)
	}

	return nil
}

func (db DB) QueryIds(s string, args ...interface{}) ([]int64, error) {
	// Prepare the SQL statement
	stmt, err := db.Conn.Prepare(s)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the row
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}

	ids := make([]int64, 0)
	for rows.Next() {
		id := int64(0)
		// Scan the row data into the i
		err = rows.Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("record not found")
			}
			return nil, fmt.Errorf("failed to retrieve row: %v", err)
		}
		ids = append(ids, id)
	}

	return ids, rows.Err()
}

func (db DB) QueryRow(s string, id int64, i ...interface{}) error {
	// Prepare the SQL statement
	stmt, err := db.Conn.Prepare(s)
	if err != nil {
		return fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the row
	row := stmt.QueryRow(id)

	// Scan the row data into the i
	err = row.Scan(i...)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("record not found")
		}
		return fmt.Errorf("failed to retrieve row: %v", err)
	}

	return nil
}

// NewDB creates a new MySQL database connection
func NewDB() (Iface, error) {
	// Replace the following values with your MySQL database credentials
	dsn := os.Getenv("db_address")

	// Create a new database connection
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check the database connection
	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Return the database connection
	return &DB{Conn: conn}, nil
}
