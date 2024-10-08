package db

import (
	"database/sql"
	"fmt"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/metric"
	"log"
)

type DB interface {
	Close() error
	SaveMetric(m metric.Metric) error
}

type dbImpl struct {
	db *sql.DB
}

func New(path string) (DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		return nil, err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS metrics (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        metric_id TEXT NOT NULL,
        name TEXT,
        uri TEXT,
        client TEXT,
        delta REAL,
        value REAL,
        rating TEXT,
        attribution TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("Failed to create table: %s", err)
	}

	fmt.Println("Database is ready, path:", path)
	return &dbImpl{db: db}, nil
}

func (d *dbImpl) SaveMetric(m metric.Metric) error {
	insertUpdateSQL := `INSERT INTO metrics (metric_id, name, uri, client, value, delta, attribution, rating) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := d.db.Exec(
		insertUpdateSQL,
		m.GetID(),
		m.GetName(),
		m.GetUri(),
		m.GetClient(),
		m.GetValue(),
		m.GetDelta(),
		m.GetAttribution(),
		m.GetRating()); err != nil {
		return fmt.Errorf("Failed to insert/update record: %s", err)
	}

	return nil
}

func (d *dbImpl) Close() error {
	return d.db.Close()
}
