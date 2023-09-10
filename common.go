package goopenai

import (
	"fmt"
)

// Error is the error standard response from the API
type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   interface{} `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

// Message is the message struct for the chat and completions endpoint
type Message struct {
	Role         string        `json:"role,omitempty"`
	Content      string        `json:"content,omitempty"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}
