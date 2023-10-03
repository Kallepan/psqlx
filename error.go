package psqlx

import "fmt"

// https://www.postgresql.org/docs/current/protocol-error-fields.html
type driverError struct {
	Severity string
	Code     string
	Message  string
}

func (d driverError) Error() string {
	return fmt.Sprintf("%s (%s): %s", d.Severity, d.Code, d.Message)
}
