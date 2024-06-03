package errors

import "encoding/json"

type CodeError struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func (error *CodeError) Error() string {
	j, err := json.MarshalIndent(error, "", "")
	if err != nil {
		panic(err)
	}
	return string(j)
}
