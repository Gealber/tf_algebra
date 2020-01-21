package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

//NewServer configure and returns a Server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	questionHandlerFunc := http.HandlerFunc(questionHandler)
	topicHandlerFunc := http.HandlerFunc(topicHandler)
	usersHandlerFunc := http.HandlerFunc(usersHandler)
	updateHandlerFunc := http.HandlerFunc(updateUserHandler)
	createHandlerFunc := http.HandlerFunc(createUser)
	mx.Handle("/api/questions", corsHandler(authHandler(questionHandlerFunc))).Methods("GET","OPTIONS")
	mx.Handle("/api/users/score", corsHandler(authHandler(updateHandlerFunc))).Methods("POST","OPTIONS")
	mx.Handle("/api/users", corsHandler(usersHandlerFunc)).Methods("GET")
	mx.Handle("/api/signUp", corsHandler(createHandlerFunc)).Methods("POST","OPTIONS")
	mx.HandleFunc("/api/login",loginHandler).Methods("POST")
	mx.Handle("/api/topic/{name}", corsHandler(authHandler(topicHandlerFunc))).Methods("GET","OPTIONS")
}

