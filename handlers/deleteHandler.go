// Package handlers handles all the GET request and responses send to the webserver.
package handlers

import (
	"log"
	"net/http"

	"../repository"
)

// HandDelete handles the /laptops GET request.
// In doing so it also parses the values from laptops.json file.
func HandDelete(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		log.Println("Deleting Comment")
		request.ParseForm()

		comments, err := repository.LoadComments()
		if err != nil {
			log.Println(err)
			http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
			return
		}
		for _, commentTD := range comments {
			if commentTD.Time == request.FormValue("time") {
				err2 := repository.DeleteComment(commentTD.Time)
				if err2 != nil {
					log.Println(err2)
					http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
					return
				}

				http.Redirect(response, request, "/", http.StatusSeeOther)
				return
			}
		}

	}
	http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
}
