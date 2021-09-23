package controllers

import (
	"flag"
	"login/config/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) InitializeRoutes(db *gorm.DB, rt *mux.Router) {
	a.DB = db
	a.Router = rt

	var dir string
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json
	flag.StringVar(&dir, "img", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	a.Router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))

	a.Router.HandleFunc("/login", a.Login).Methods("POST")
	a.Router.HandleFunc("/register", a.Register).Methods("POST")

	s := a.Router.PathPrefix("/api").Subrouter() // subrouter to add auth middleware
	s.Use(middlewares.SetContentTypeMiddleware)  // setting content-type to json
	s.Use(middlewares.AuthJwtVerify)

}
