package repository

import (
	"log"

	. "../types"
)

const (
	COMMENTFILE = "./data/comment.json"
	TAGFILE     = "./data/tags.json"
)

// GetLaptop takes the index number of a laptop and compares that to that of the

// LoadComments loads the data from the the laptops.json file into a slice from the Laptop data struct, exporting it.
func LoadComments() ([]Comment, error) {
	var comments []Comment
	err := LoadData(COMMENTFILE, &comments)

	return comments, err
}

func LoadTags() ([]Tag, error) {
	var tags []Tag
	err := LoadData(TAGFILE, &tags)

	return tags, err
}

// SaveComment checks whether the customer added already exists and if not adds the new customer to the list and saves it to customers.json.
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

func SaveTag(data string) error {
	log.Println("Saving Tag")

	tags, err := LoadTags()
	if err != nil {
		log.Println(err)
		return err
	}

	tag := Tag{
		Index: len(tags),
		Tag:   data,
	}
	tags = append(tags, tag)

	return SaveData(TAGFILE, tags)
}

// DeleteReview takes the time from a review in the from of a string and deletes.
func DeleteComment(date string) error {
	log.Println("Deleting Comment")
	comments, err := LoadComments()
	if err != nil {
		return err
	}

	for i, comment := range comments {
		if comment.Time == date {
			comments = append(comments[:i], comments[deleteCheck(i, len(comments)):]...)
		}
	}

	return SaveData(COMMENTFILE, comments)
}

func deleteCheck(value int, length int) int {
	if value+1 > length {
		return length
	}

	return value + 1
}

type ByIndex []Tag

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].Index > a[j].Index }
