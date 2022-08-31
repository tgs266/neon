package errors

type SerializableError struct {
	Message    string         `json:"message"`
	ErrorCode  string         `json:"errorCode"`
	Parameters map[string]any `json:"parameters"`
}
