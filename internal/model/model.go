package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/codenotary/immudb/pkg/stdlib"
)

type Log struct {
	Id       int       `json:"id"`
	DateTime time.Time `json:"datetime"`
	Data     string    `json:"data"`
	Source   string    `json:"source"`
}

type MyLogRepository interface {
	InsertLog(log *Log) error
	FetchLog(lastX int) ([]*Log, error)
	CreateLogTable() error
	Close() error
}

type MyLogImmudbRepo struct {
	ctx context.Context
	db  *sql.DB
}

// NewMyLogImmudbRepo constructs MyLogImmudbRepo object.
func NewMyLogImmudbRepo(ctx context.Context, dsn string) (MyLogRepository, error) {
	db, err := getConnection(dsn)
	if err != nil {
		return nil, err
	}
	return MyLogImmudbRepo{
		ctx: ctx,
		db:  db,
	}, nil
}

func (m MyLogImmudbRepo) FetchLog(lastX int) ([]*Log, error) {
	q := "SELECT * FROM myLog;"
	if lastX != 0 {
		q = fmt.Sprintf("SELECT * FROM myLog ORDER BY id DESC LIMIT %v;", lastX)
	}
	res, err := m.db.QueryContext(m.ctx, q)
	if err != nil {
		return nil, fmt.Errorf("Query err: %w", err)
	}
	result := make([]*Log, 0)
	for res.Next() {
		lr := Log{}
		err = res.Scan(&lr.Id, &lr.DateTime, &lr.Data, &lr.Source)
		if err != nil {
			return nil, err //nolint:wrapcheck
		}
		result = append(result, &lr)
	}
	return result, nil
}
func (m MyLogImmudbRepo) InsertLog(log *Log) error {
	q := fmt.Sprintf("INSERT INTO myLog ( datetime, data, source) VALUES ( NOW(), '%v' , '%v');",
		log.Data, log.Source)
	_, err := m.db.ExecContext(m.ctx, q)
	if err != nil {
		return fmt.Errorf("Query err: %w", err)
	}
	return nil
}

func (m MyLogImmudbRepo) CreateLogTable() error {
	q := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS myLog(id INTEGER AUTO_INCREMENT, datetime TIMESTAMP, data VARCHAR, source VARCHAR, PRIMARY KEY id)")
	_, err := m.db.ExecContext(m.ctx, q)
	if err != nil {
		return fmt.Errorf("Query err: %w", err)
	}
	return nil
}

func (m MyLogImmudbRepo) Close() error {
	return m.db.Close()
}

func getConnection(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("immudb", dsn)
	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, err
	}
	return db, err
}
