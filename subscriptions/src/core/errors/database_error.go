package errors

import (
	"encoding/json"
)

type DatabaseError struct {
	CodeError `json:"CodeError"`
}

func (error *DatabaseError) Error() string {
	j, err := json.MarshalIndent(error, "", "")
	if err != nil {
		panic(err)
	}
	return string(j)
}
