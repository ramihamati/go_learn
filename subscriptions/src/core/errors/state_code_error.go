package errors

import "encoding/json"

// StateCodeErrorDetails extra details for an error
type StateCodeErrorDetails struct {
	// HasAssociatedKey if the error represents a validation error or is
	// related to certain operation the key would hold a name reference to that
	HasAssociatedKey string `json:"HasAssociatedKey"`
	AssociatedKey    string `json:"AssociatedKey"`
}

// StateCodeFailure the error which is represented by the code and a message
type StateCodeFailure struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`

	// HasObjectResult is the error has an extra object as detail
	// the objectResult would store that and the typename is used for
	// deserialization
	HasObjectResult      bool        `json:"HasObjectResult"`
	ObjectResultTypeName string      `json:"ObjectResultTypeName"`
	ObjectResult         interface{} `json:"ObjectResult"`
}

type StateCodeEntry struct {
	Details  StateCodeErrorDetails `json:"Details"`
	Failures []StateCodeFailure    `json:"Failures"`
}

func (entry *StateCodeEntry) IsErrored() bool {
	return len(entry.Failures) > 0
}

func (entry *StateCodeEntry) IsSuccessful() bool {
	return len(entry.Failures) == 0
}

type StateCodeError struct {
	Status  uint16           `json:"Status"`
	Entries []StateCodeEntry `json:"Entries"`
}

func (error *StateCodeError) Error() string {
	j, err := json.MarshalIndent(error, "", "")
	if err != nil {
		panic(err)
	}

	return string(j)
}
