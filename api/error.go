package vbuy

import (
    "fmt"
    "net/http"
)

type Error struct {
    Code   string
    Title  string
    Status int
    Detail []*ErrorDetail
}

type ErrorDetail struct {
    Field    string   `json:"field"`
    Messages []string `json:"messages"`
}

func (ed *ErrorDetail) AppendMessage(m string) {
    ed.Messages = append(ed.Messages, m)
}

func (ed *ErrorDetail) HasMessages() bool {
    return ed.Messages != nil && len(ed.Messages) > 0
}

type ErrorResponse struct {
    Code   string         `json:"code"`
    Title  string         `json:"title"`
    Detail []*ErrorDetail `json:"detail,omitempty"`
}

func (e Error) Error() string {
    return fmt.Sprintf("Code: %s, Title: %s, Status: %d", e.Code, e.Title, e.Status)
}

func (e Error) ToErrorResponse() *ErrorResponse {
    return &ErrorResponse{
        Code:   e.Code,
        Title:  e.Title,
        Detail: e.Detail,
    }
}

var (
    // Common errors
    ErrInternalServer = Error{Code: "Vb001", Title: "Server failure", Status: http.StatusInternalServerError}
    ErrUnauthorized   = Error{Code: "Vb002", Title: "Not authorized to perform the action", Status: http.StatusUnauthorized}
    ErrNotImplemented = Error{Code: "Vb003", Title: "Action not implemented", Status: http.StatusNotImplemented}
    ErrInvalidPayload = Error{Code: "Vb004", Title: "Payload is invalid", Status: http.StatusBadRequest}

    // Products
    ErrProductNotFound      = Error{Code: "VbProduct001", Title: "Product not found", Status: http.StatusNotFound}
    ErrProductAlreadyExists = Error{Code: "VbProduct002", Title: "Product already exists", Status: http.StatusBadRequest}
    ErrProductIsNotActive   = Error{Code: "VbProduct003", Title: "Product is not active", Status: http.StatusBadRequest}
)
