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
        id TEXT NOT NULL PRIMARY KEY,
        name TEXT,
        uri TEXT,
        client TEXT,
        value REAL,
        target TEXT,
        rating TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("Failed to create table: %s", err)
	}

	fmt.Println("Database is ready, path:", path)
	return &dbImpl{db: db}, nil
}

func (d *dbImpl) SaveMetric(m metric.Metric) error {
	insertUpdateSQL := `INSERT INTO metrics (id, name, uri, client, value, target, rating) VALUES (?, ?, ?, ?, ?, ?, ?)
                        ON CONFLICT(id) DO UPDATE SET
                        name = excluded.name,
						uri = excluded.uri,
						client = excluded.client,
                        value = excluded.value,
                        target = excluded.target,
                        rating = excluded.rating;`
	if _, err := d.db.Exec(insertUpdateSQL, m.GetID(), m.GetName(), m.GetUri(), m.GetClient(), m.GetValue(), m.GetTarget(), m.GetRating()); err != nil {
		return fmt.Errorf("Failed to insert/update record: %s", err)
	}

	return nil
}

func (d *dbImpl) Close() error {
	return d.db.Close()
}
