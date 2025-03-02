package database

import (
	"database/sql"
	"sync"

	"github.com/ville-koskela/go-ldap-server/domain"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite3Database struct {
	db *sql.DB
	mu sync.Mutex
}

func NewSQLite3Database(dataSourceName string) (*SQLite3Database, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Create table if not-exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (username TEXT PRIMARY KEY, password TEXT, email TEXT, full_name TEXT, uid INTEGER, gid INTEGER)")

	if err != nil {
		return nil, err
	}

	return &SQLite3Database{db: db}, nil
}

func (db *SQLite3Database) AddUser(user domain.User) (domain.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, err := db.db.Exec("INSERT INTO users (username, password, email, full_name, uid, gid) VALUES (?, ?, ?, ?, ?, ?)",
		user.Username, user.Password, user.Email, user.FullName, user.UID, user.GID)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (db *SQLite3Database) FindUserByUsername(username string) (domain.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	row := db.db.QueryRow("SELECT username, password, email, full_name, uid, gid FROM users WHERE username = ?", username)

	var user domain.User
	err := row.Scan(&user.Username, &user.Password, &user.Email, &user.FullName, &user.UID, &user.GID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (db *SQLite3Database) ListUsers() ([]domain.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	rows, err := db.db.Query("SELECT username, password, email, full_name, uid, gid FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Username, &user.Password, &user.Email, &user.FullName, &user.UID, &user.GID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (db *SQLite3Database) Close() error {
	return db.db.Close()
}
