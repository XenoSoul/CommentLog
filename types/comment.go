package types

import (
	"time"
)

// Comment is exclusively used for the admin functions and contains whatever values that need to be saved for futher use.
// Primarily this struct is used to keep track of those who have been banned and the words that are stored for the filtering of comments.
type Comment struct {
	Index   int       `json:"index"`
	Comment string    `json:"comment"`
	Time    time.Time `json:"time"`
}
