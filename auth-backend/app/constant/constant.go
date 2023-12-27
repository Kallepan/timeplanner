package constant

type ResponseStatus int
type Headers int
type General int

// Constant API
const (
	Success ResponseStatus = iota + 1
	InvalidRequest
	Unauthorized
	DataNotFound
	UnknownError
	DatabaseError
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{
		"Success",
		"Invalid Request",
		"Unauthorized",
		"Data Not Found",
		"Unknown Error",
		"Database Error",
	}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return  [...]string{
		"Success",
		"Invalid Request: Please check your request",
		"Unauthorized: Please check your credentials",
		"Data Not Found: Data not found",
		"Unknown Error: Unknown error",
		"Database Error: Error when executing query to database",
	}[r-1]
}