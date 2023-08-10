package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"remote-task/utilities"
	"sync"
	"time"
)

type Repositories struct {
	Db         *sql.DB
	Statements map[string]*sql.Stmt
	mu         sync.Mutex
}

const Dialect = "mysql"

// NewRepositories create new my sql instance for other repositories
func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%v)/%s?collation=utf8mb4_unicode_ci&parseTime=True",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
	db, err := sql.Open(Dialect, dsn)
	if err != nil {
		fmt.Println("db connection failed", err)
	}
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(2)

	if pingErr := db.Ping(); pingErr != nil {
		log.Println("Err mysql ping", pingErr)
	} else {
		log.Println("Success mysql connection is ok")
	}

	var skip string
	var maxConnections int
	maxConErr := db.QueryRow(utilities.SHOW_VARS_CONNECTION).Scan(&skip, &maxConnections)
	if maxConErr != nil {
		log.Println("Err mysql getting the max_connections", maxConErr)
	}
	maxConnections = int(math.Floor(float64(maxConnections) * 0.9))
	if maxConnections == 0 {
		maxConnections = 100
	}
	db.SetMaxOpenConns(maxConnections)

	var waitTimeout int
	waitErr := db.QueryRow(utilities.SHOW_VARS_TIMEOUT).Scan(&skip, &waitTimeout)
	if waitErr != nil {
		log.Println("Err mysql getting the wait_timeout", waitErr)
	}
	if waitTimeout == 0 {
		waitTimeout = 180
	}
	waitTimeout = int(math.Min(float64(waitTimeout), 180))
	t := time.Duration(waitTimeout) * time.Second
	db.SetConnMaxLifetime(t)

	return &Repositories{
		Db:         db,
		Statements: make(map[string]*sql.Stmt),
	}, nil
}

// Stmt declare stmt the database connection
func (mr *Repositories) Stmt(id string) *sql.Stmt {
	return mr.Statements[id]
}

// SetStmt set stmt the database connection
func (mr *Repositories) SetStmt(id string, stmt *sql.Stmt) {
	mr.Statements[id] = stmt
}

// Ping pings the database connection
func (mr *Repositories) Ping() error {
	if err := mr.Db.Ping(); err != nil {
		return err
	}

	return nil
}

// DB returns the database connection
func (mr *Repositories) DB() *sql.DB {
	return mr.Db
}

// Close closes the  database connection
func (mr *Repositories) Close() {
	for _, stmt := range mr.Statements {
		_ = stmt.Close()
	}
	_ = mr.Db.Close()
}
