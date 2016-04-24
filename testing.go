package que

import "github.com/jackc/pgx"

// TestInjectJobConn inject *pgx.Conn to Job
func TestInjectJobConn(j *Job, conn *pgx.Conn) *Job {
	j.conn = conn
	return j
}
