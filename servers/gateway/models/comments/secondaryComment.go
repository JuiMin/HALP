package comments

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SecondaryComment stores the secondary comments, which are responses to comments
// allowing further discussion
type SecondaryComment struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	ImageURL    string        `json:"image_url"`
	Content     string        `json:"content"`
	Parent      bson.ObjectId `json:"parent"`
	AuthorID    bson.ObjectId `json:"author_id"`
	PostID      bson.ObjectId `json:"post_id"`
	Upvotes     int           `json:"upvotes"`
	Downvotes   int           `json:"downvotes"`
	TotalVotes  int           `json:"total_votes"`
	TimeCreated time.Time     `json:"time_created"`
	TimeEdited  time.Time     `json:"time_edited"`
	Secondary   bool          `json:"secondary"`
}

// NewSecondaryComment contains the information required for a secondary comment
type NewSecondaryComment struct {
	ImageURL string        `json:"image_url"`
	Content  string        `json:"content"`
	PostID   bson.ObjectId `json:"post_id"`
	Parent   bson.ObjectId `json:"parent"`
	AuthorID bson.ObjectId `json:"author_id"`
}

// SecondaryCommentUpdate contains all the information that could be updated in a comment
type SecondaryCommentUpdate struct {
	ImageURL string `json:"image_url"`
	Content  string `json:"content"`
}

// SecondaryCommentVote contains an integer that represents the vote of this user
type SecondaryCommentVote struct {
	Upvote   int `json:"upvote"`
	Downvote int `json:"downvote"`
}

// Validate should validate the new comment object to confirm that we have a proper comment
func (nc *NewSecondaryComment) Validate() error {
	// Check that the comment contains meaningful information
	if len(nc.ImageURL) == 0 && len(nc.Content) == 0 {
		return fmt.Errorf("Cannot have no image or content in comment")
	}

	if !bson.IsObjectIdHex(nc.AuthorID.Hex()) {
		return fmt.Errorf("Error: Invalid Author ID")
	}

	if !bson.IsObjectIdHex(nc.PostID.Hex()) {
		return fmt.Errorf("Error: Invalid Post ID")
	}

	if !bson.IsObjectIdHex(nc.Parent.Hex()) {
		return fmt.Errorf("Error: Invalid Parent ID")
	}
	// Passed the validation
	return nil
}

// ToSecondaryComment takes a new comment and converts it to a comment object
func (nc *NewSecondaryComment) ToSecondaryComment() (*SecondaryComment, error) {

	// Validate the new comment structure is admissable
	if err := nc.Validate(); err != nil {
		return nil, err
	}

	// Construct the new comment
	comment := &SecondaryComment{
		ID:          bson.NewObjectId(),
		ImageURL:    nc.ImageURL,
		Content:     nc.Content,
		AuthorID:    nc.AuthorID,
		Parent:      nc.Parent,
		PostID:      nc.PostID,
		Upvotes:     0,
		Downvotes:   0,
		TotalVotes:  0,
		TimeCreated: time.Now(),
		TimeEdited:  time.Now(),
		Secondary:   true,
	}

	// Return the created comment
	return comment, nil
}

// Update alters the composition of a comment based on the attributes in the update struct
// The alterable components are changed here.
func (c *SecondaryComment) Update(updates *SecondaryCommentUpdate) error {
	// Check for valid updates
	if len(updates.Content) == 0 && len(updates.ImageURL) == 0 {
		return fmt.Errorf("We cannot set the comment to contain nothing")
	}

	// Valid updates
	c.ImageURL = updates.ImageURL
	c.Content = updates.Content
	// Update the time stamps
	c.TimeEdited = time.Now()

	// No errors to report
	return nil
}

// Vote allows for the shifting of the votes on a comment
func (c *SecondaryComment) Vote(v *SecondaryCommentVote) error {
	// Alter the votes based on the input from the update

	if c.Upvotes+v.Upvote <= 1 && c.Upvotes+v.Upvote >= -1 {
		if c.Downvotes+v.Downvote <= 1 && c.Downvotes+v.Downvote >= -1 {

			if v.Upvote == 1 && v.Downvote == 1 {
				return fmt.Errorf("Can't increment upvote and downvote")
			}

			if v.Upvote == -1 && v.Downvote == -1 {
				return fmt.Errorf("Can't decrement both upvote and downvote")
			}

			c.Upvotes += v.Upvote
			c.Downvotes += v.Downvote

			// Don't allow negatives
			if c.Upvotes < 0 {
				c.Upvotes = 0
			}

			if c.Downvotes < 0 {
				c.Downvotes = 0
			}
		}
		return nil
	}

	return fmt.Errorf("Invalid Votes Received, aborting changes")
}
