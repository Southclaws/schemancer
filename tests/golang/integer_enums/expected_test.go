package integer_enums_test

type HttpStatus int

const (
	HttpStatus200 HttpStatus = 200
	HttpStatus201 HttpStatus = 201
	HttpStatus400 HttpStatus = 400
	HttpStatus404 HttpStatus = 404
	HttpStatus500 HttpStatus = 500
)

type Priority int

const (
	Priority1 Priority = 1
	Priority2 Priority = 2
	Priority3 Priority = 3
)

type Response struct {
	Priority *Priority  `json:"priority,omitempty"`
	Status   HttpStatus `json:"status"`
}
