package model

const CREATE_SUBMISSION_TABLE = `CREATE TABLE IF NOT EXISTS submission (
	sub_id INTEGER PRIMARY KEY AUTOINCREMENT, 
	event_id INTEGER NOT NULL,
	group_id INTEGER NOT NULL,
	problem_id INTEGER NOT NULL,
	source_code TEXT,
	language TEXT NOT NULL,
	status INTEGER NOT NULL DEFAULT 0,
	accuracy REAL DEFAULT 0.0,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT DEFAULT CURRENT_TIMESTAMP);`

const CREATE_SUBMISSION_INDEX1 = `CREATE INDEX IF NOT EXISTS submission_event_status_idx ON submission (event_id,status);`
const CREATE_SUBMISSION_INDEX2 = `CREATE INDEX IF NOT EXISTS submission_probid_groupid_idx ON submission (problem_id,group_id);`

const INSERT_SUBMISSION = `INSERT INTO submission (event_id, group_id, problem_id, source_code, language, status) VALUES (?, ?, ?, ?, ?, ?)`

const UPDATE_SUBMISSION = `UPDATE submission SET updated_at = DATETIME('NOW'), status = ?, accuracy = ? WHERE sub_id = ?`

const SELECT_SUBMISSION = `SELECT sub_id, event_id, group_id, problem_id, source_code, language, status, accuracy, created_at, updated_at FROM submission WHERE event_id = ? AND group_id = ? LIMIT ? OFFSET ?`

const SELECT_LEADERS = `SELECT group_id, SUM(submission_processed.max_accuracy) AS score
	FROM   ( SELECT   group_id, Max(accuracy) AS max_accuracy 
             FROM submission 
             WHERE event_id = ? AND status = ?
			 GROUP BY group_id, problem_id
	) AS submission_processed 
	GROUP BY submission_processed.group_id 
	ORDER BY score DESC LIMIT ?`

const FIND_ALL_TOP_SOLVED_PROBLEMS = `SELECT sub_id, problem_id, language, Max(accuracy) AS max_accuracy, status, updated_at
             FROM submission 
             WHERE event_id = ? AND group_id = ? AND status = ?
			 GROUP BY problem_id 
			 ORDER BY max_accuracy DESC LIMIT ?`
