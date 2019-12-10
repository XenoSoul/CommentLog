package repository

import (
	"log"

	. "../types"
)

const (
	COMMENTFILE = "./data/comment.json"
)

// GetLaptop takes the index number of a laptop and compares that to that of the

// LoadComments loads the data from the the laptops.json file into a slice from the Laptop data struct, exporting it.
func LoadComments() ([]Comment, error) {
	var comments []Comment
	err := LoadData(COMMENTFILE, &comments)

	return comments, err
}

// SaveAdmin checks whether the customer added already exists and if not adds the new customer to the list and saves it to customers.json.
func SaveComment(data Comment) error {
	log.Println("Saving Comment")

	comments, err := LoadComments()
	if err != nil {
		log.Println(err)
		return err
	}
	comments = append(comments, data)

	return SaveData(COMMENTFILE, comments)
}

// DeleteReview takes the time from a review in the from of a string and deletes.
// func DeleteReview(date string) error {
// 	log.Println("Deleting Review")
// 	reviews, err := LoadReviews()
// 	if err != nil {
// 		return err
// 	}

// 	for i, review := range reviews {
// 		if (review.Date).String() == date {
// 			reviews = append(reviews[:i], reviews[deleteCheck(i, len(reviews)):]...)
// 		}
// 	}

// 	return SaveData(REVIEWSFILE, reviews)
// }

// func deleteCheck(value int, length int) int {
// 	if value+1 > length {
// 		return length
// 	}

// 	return value + 1
// }
