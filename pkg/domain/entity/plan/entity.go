package planentity

import "time"

type Reserve struct {
	ID        string
	StartedAt time.Time
}

type Status string

const (
	PlanStatusOpen  = "open"
	PlanStatusClose = "close"
)

type Plan struct {
	ID       string
	Name     string
	Status   Status
	Reserves []*Reserve
}
