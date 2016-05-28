package que

// TestInjectJobConn injects *pgx.Conn to Job
func TestInjectJobConn(j *Job, conn Ext) *Job {
	j.conn = conn
	return j
}

// TestInjectJobTx injects tx to Job
func TestInjectJobTx(j *Job, tx Txer) *Job {
	j.tx = tx
	return j
}
