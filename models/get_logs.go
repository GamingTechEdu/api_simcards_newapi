package models

type GetLogs struct {
	Logid     string `json:"log_id"`
	SimcardId string `json:"simcard_id"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
