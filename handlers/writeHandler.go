// Package handlers handles all the GET request and responses send to the webserver.
package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"../repository"
	. "../types"
)

// HandComment handles the /laptops GET request.
// In doing so it also parses the values from laptops.json file.
func HandComment(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		log.Println("Savings comments")
		request.ParseForm()

		comments, err3 := repository.LoadComments()
		if err3 != nil {
			log.Println(err3)
			http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
			return
		}

		comment := Comment{
			Index:   len(comments),
			Time:    time.Now(),
			Comment: request.FormValue("comment"),
		}

		err2 := repository.SaveComment(comment)
		if err2 != nil {
			log.Println(err2)
			http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
			return
		}
	}
	log.Println("Viewing comments")

	comments, err := repository.LoadComments()
	if err != nil {
		log.Println(err)
		http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
		return
	}

	// template.ParseFiles parses the wanted file to be rendered and send back to the user.
	render, err1 := template.ParseFiles("templates/comments.html")
	if err1 != nil {
		log.Println(err1)
		http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
		return
	}

	// render.Execute sends a response to the user while parsing the laptops variables to the webpage.
	render.Execute(response, comments)
}
