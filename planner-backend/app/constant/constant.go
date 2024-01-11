package constant

type ResponseStatus int
type Headers int
type General int

// App Constant
const (
	TimeFormat string = "15:04"
)

// Constant API
const (
	Success ResponseStatus = iota + 1
	InvalidRequest
	Unauthorized
	DataNotFound
	Conflict
	UnknownError
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{
		"Success",
		"Invalid Request",
		"Unauthorized",
		"Data Not Found",
		"Conflict",
		"Unknown Error",
	}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{
		"Success",
		"Invalid Request: Please check your request",
		"Unauthorized: Please check your credentials",
		"Data Not Found: Data not found",
		"Conflict: Data already exist",
		"Unknown Error: Unknown error",
	}[r-1]
}
