package valueobjects

type AssignmentStatus string

const (
	NotStarted AssignmentStatus     = "NotStarted"
	Started AssignmentStatus        = "Started"
	ReadyForReview AssignmentStatus = "ReadyForReview"
	Completed AssignmentStatus      = "Completed"
	Canceled AssignmentStatus       = "Canceled"
)