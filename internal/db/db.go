package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Database struct to hold multiple connections
type Database struct {
	FIS       *sql.DB
	UTV       *sql.DB
	Auth      *sql.DB
	Tietoevry *sql.DB
	KAMK      *sql.DB
}

func NewSingleDB(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := connectDB(addr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func New(fisAddr, utvAddr, authAddr, tietoevryAddr, kamkAddr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*Database, error) {
	fisDB, err := connectDB(fisAddr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	utvDB, err := connectDB(utvAddr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	authDB, err := connectDB(authAddr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	tietoevryDB, err := connectDB(tietoevryAddr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	kamkDB, err := connectDB(kamkAddr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	return &Database{FIS: fisDB, UTV: utvDB, Auth: authDB, Tietoevry: tietoevryDB, KAMK: kamkDB}, nil
}

func connectDB(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(time.Duration(duration))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
