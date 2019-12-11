package types

// Tag is exclusively used for the admin functions and contains whatever values that need to be saved for futher use.
// Primarily this struct is used to keep track of those who have been banned and the words that are stored for the filtering of comments.
type Tag struct {
	Index int    `json:"index"`
	Tag   string `json:"tag"`
}
