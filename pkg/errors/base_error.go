package errors

// BaseErrorResponse represents the standard error response structure
type BaseErrorResponse struct {
	Status         bool                   `json:"status"`
	Code           int                    `json:"code"`
	Message        string                 `json:"message"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo,omitempty"`
}
