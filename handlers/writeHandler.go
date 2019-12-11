// Package handlers handles all the GET request and responses send to the webserver.
package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"../repository"
	. "../types"
)

type commentData struct {
	Tag     []Tag
	Comment []Comment
	Recent  string
}

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

		var tag string
		if request.FormValue("tag") != request.FormValue("tagn") {
			tag = request.FormValue("tagn")
			err := repository.SaveTag(tag)
			if err != nil {
				log.Println(err)
				http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
				return
			}
		} else {
			tag = request.FormValue("tag")
		}

		var commentSt string
		if request.FormValue("commenth") == "" {
			commentSt = request.FormValue("commentb")
		} else {
			commentSt = request.FormValue("commenth") + "\n" + request.FormValue("commentb")
		}

		comment := Comment{
			Index:   len(comments),
			Tag:     tag,
			Time:    (time.Now()).Format(time.RFC1123),
			Comment: commentSt,
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
	sort.Sort(ByIndex(comments))

	tags, err2 := repository.LoadTags()
	if err2 != nil {
		log.Println(err2)
		http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
		return
	}

	sort.Sort(repository.ByIndex(tags))
	var recent string
	if len(tags) == 0 {
		recent = ""
	} else {
		recent = comments[0].Tag
	}

	data := commentData{
		Comment: comments,
		Tag:     tags,
		Recent:  recent,
	}

	// template.ParseFiles parses the wanted file to be rendered and send back to the user.
	render, err1 := template.ParseFiles("templates/comments.html")
	if err1 != nil {
		log.Println(err1)
		http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
		return
	}

	// render.Execute sends a response to the user while parsing the laptops variables to the webpage.
	render.Execute(response, data)
}

type ByIndex []Comment

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].Index > a[j].Index }
