package enums

type HttpMethod string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodPatch   HttpMethod = "PATCH"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodOptions HttpMethod = "OPTIONS"
)

type ApiRequest struct {
	Body   *string    `json:"body,omitempty"`
	Method HttpMethod `json:"method"`
	URL    string     `json:"url"`
}

type Color string

const (
	ColorRed    Color = "red"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorYellow Color = "yellow"
)

type Priority string

const (
	PriorityLow      Priority = "low"
	PriorityMedium   Priority = "medium"
	PriorityHigh     Priority = "high"
	PriorityCritical Priority = "critical"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
	StatusCancelled  Status = "cancelled"
)

type Task struct {
	Color    *Color    `json:"color,omitempty"`
	ID       string    `json:"id"`
	Priority *Priority `json:"priority,omitempty"`
	Status   Status    `json:"status"`
	Title    string    `json:"title"`
}
