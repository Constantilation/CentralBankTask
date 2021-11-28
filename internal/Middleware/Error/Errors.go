package Error

// MultiLogger implementation of logger interface methods
type MultiLogger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Sync() error
}

// ResultError response error structure
type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

// Errors another version of errors struct
type Errors struct {
	Alias string
	Text  string
}

// Error method to return Alias of error
func (e *Errors) Error() string {
	return e.Alias
}

// CheckError struct with MultiLogger, can add another loggers or smth else
type CheckError struct {
	RequestId int
	Logger    MultiLogger
}

// Error of server
const (
	ErrAtoi            = "func Atoi convert string in int"
	ErrCreate          = "Cannot create file"
	ErrGet             = "Cannot get data from server"
	ErrCopy            = "Cannot copy data to file"
	ErrNotStringAndInt = "expected type string or int"
	IntNil             = 0
)

// Error of main
const (
	MCreateDBNotConnect = "db not connect"
)

// Transaction errors
const (
	UpdateTransactionNotCreated = "UpdateTransactionNotCreated"
	ValutesNotInsert            = "ValutesNotInsert"
	ValutesNotCommit            = "ValutesNotCommit"
	DateNotInserter             = "DateNotInserter"
	DateNotCommit               = "DateNotCommit"
)

// Date check errors
const (
	NotValidDate = "NotValidDate"
)
