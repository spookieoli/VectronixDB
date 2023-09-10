package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// HttpHandler handles http requests
type HttpHandler struct {
	Port      int
	Server    http.Server
	templates *template.Template
}

// NewHttpHandler creates a new HttpHandler
func NewHttpHandler(port int) *HttpHandler {
	httpHandler := &HttpHandler{
		Port: port,
		Server: http.Server{
			Addr: ":" + strconv.Itoa(port),
		},
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}
	return httpHandler
}

// connectRoutes connects the routes to the http handler
func (h *HttpHandler) connectRoutes() {
	http.HandleFunc("/vectronix/index", h.Home)
}

// Start starts the http Server
func (h *HttpHandler) Start() {
	err := h.Server.ListenAndServe()
	if err != nil {
		// Everything for now :)
		panic(err)
	}
}

/* Connect the Routes */
func (h *HttpHandler) Home(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		// TODO: Handle error + Create Logger
		fmt.Println(err)
	}
}
