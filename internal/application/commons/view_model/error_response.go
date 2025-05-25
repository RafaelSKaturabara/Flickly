package view_model

type ErrorResponse struct {
	Code            int    `json:"code"`
	Message         string `json:"message"`
	InternalMessage string `json:"internalMessage,omitempty"`
}
