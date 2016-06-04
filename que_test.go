package que

import (
	"testing"

	"github.com/jackc/pgx"
)

var testConnConfig = pgx.ConnConfig{
	Host:     "localhost",
	Database: "que-go-test",
}

func openTestClientMaxConns(t testing.TB, maxConnections int) *Client {
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     testConnConfig,
		MaxConnections: maxConnections,
		AfterConnect:   PrepareStatements,
	}
	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		t.Fatal(err)
	}
	return NewClient(pool)
}

func openTestClient(t testing.TB) *Client {
	return openTestClientMaxConns(t, 5)
}

func truncateAndClose(pool *pgx.ConnPool) {
	if _, err := pool.Exec("TRUNCATE TABLE que_jobs"); err != nil {
		panic(err)
	}
	pool.Close()
}

type queriable interface {
	Query(query string, args ...interface{}) (*pgx.Rows, error)
	QueryRow(query string, args ...interface{}) *pgx.Row
	Exec(query string, args ...interface{}) (pgx.CommandTag, error)
}

func findOneJob(q queriable) (*Job, error) {
	findSQL := `
	SELECT priority, run_at, job_id, job_class, args, error_count, last_error, queue
	FROM que_jobs LIMIT 1`

	j := &Job{}
	err := q.QueryRow(findSQL).Scan(
		&j.Priority,
		&j.RunAt,
		&j.ID,
		&j.Type,
		&j.Args,
		&j.ErrorCount,
		&j.LastError,
		&j.Queue,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return j, nil
}
