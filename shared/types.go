package shared

type Status = string

const (
	StatusPending Status = "pending"
	StatusSynced  Status = "synced"
	StatusFailed  Status = "failed"
)
