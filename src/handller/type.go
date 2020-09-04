package handller

type PostSubmissionPayload struct {
	EventId    int64  `json:"event_id"`
	SourceCode string `json:"source_code"`
	Lang       string `json:"lang"`
	ProblemId  int64  `json:"problem_id"`
}
