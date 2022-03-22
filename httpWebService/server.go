package main

import "net/http"

func New() *server {
	s := &server{}
	s.routes()
	return s
}

// start your server
func (s *server) ServeHttp(w *http.ResponseWriter, r *http.Request) {
	s.router.ServeHttp(w, r)
}

//set up routes over here
func (s *server) routes() {
	s.router.Get("/api", s.handleAPI())
}

// api exmaple
func (s *server) handleAPI(w *http.ResponseWriter, r *http.Request) {
	//do something
}

// handlers
func (s *server) handleSomething() http.HandlerFunc {
	// do something
	// get things
	return func(w *http.ResponseWriter, r *http.Request) {
		//use thing
	}
}
