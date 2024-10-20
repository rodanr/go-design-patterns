package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type AuthHandler struct {
	BaseHandler
}

func (h *AuthHandler) Handle(request string) {
	if request == "auth" {
		fmt.Println("AuthHandler: Handling authentication")
	} else {
		fmt.Println("AuthHandler: Passing to next")
		h.BaseHandler.Handle(request)
	}
}

type LoggingHandler struct {
	BaseHandler
}

func (h *LoggingHandler) Handle(request string) {
	fmt.Println("LoggingHandler: Logging request")
	h.BaseHandler.Handle(request)
}

type DataHandler struct {
	BaseHandler
}

func (h *DataHandler) Handle(request string) {
	if request == "data" {
		fmt.Println("DataHandler: Handling data request")
	} else {
		fmt.Println("DataHandler: Passing to next")
		h.BaseHandler.Handle(request)
	}
}

func main() {
	authHandler := &AuthHandler{}
	loggingHandler := &LoggingHandler{}
	dataHandler := &DataHandler{}

	authHandler.SetNext(loggingHandler).SetNext(dataHandler)

	fmt.Println("Request: auth")
	authHandler.Handle("auth")
	fmt.Println()

	fmt.Println("Request: data")
	authHandler.Handle("data")
	fmt.Println()

	fmt.Println("Request: unknown")
	authHandler.Handle("unknown")
}
