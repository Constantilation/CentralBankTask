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
	ErrDB              = "database is not responding"
	ErrEncode          = "Encode"
	ErrAtoi            = "func Atoi convert string in int"
	ErrCreate          = "Cannot create file"
	ErrGet             = "Cannot get data from server"
	ErrCopy            = "Cannot copy data to file"
	ErrNotStringAndInt = "expected type string or int"
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	ErrUnmarshal       = "unmarshal json"
	IntNil             = 0
)

// Error of main
const (
	MCreateDBNotConnect         = "db not connect"
	MCreateDBCreateFileNotFound = "CreateTables.sql not found"
	MCreateDBDeleteFileNotFound = "DeleteTables.sql not found"
	MCreateDBFillFileNotFound   = "Fill.sql not found"
	MCreateDBNotCreateTables    = "table not create"
	MCreateDBNotDeleteTables    = "table not delete"
	MCreateDBNotFillTables      = "table not fill"
)

// Transaction errors
const (
	UpdateTransactionNotCreated = "UpdateTransactionNotCreated"
	ValutesNotInsert            = ""
	ValutesNotCommit            = "ValutesNotCommit"
)
