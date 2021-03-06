package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/JuiMin/HALP/servers/gateway/models/posts"
	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
)

//NewPostHandler handles requests related to Posts
//POST /posts/new
func (cr *ContextReceiver) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		errorMessage := ""
		canProceed := true
		status := http.StatusAccepted
		if r.Body == nil {
			errorMessage = "Error: Could not decode request body"
			status = http.StatusBadRequest
			canProceed = false
		}
		if canProceed {
			// Get the new post
			newPost := &posts.NewPost{}
			err := json.NewDecoder(r.Body).Decode(newPost)
			if err != nil {
				errorMessage = "Error: could not decode request body"
				status = http.StatusBadRequest
				canProceed = false
			}

			_, err = cr.BoardStore.GetByID(newPost.BoardID)
			if err != nil {
				errorMessage = "Error: board does not exist"
				status = http.StatusBadRequest
				canProceed = false
			}
			if canProceed {
				// Check for duplicate
				_, err = cr.BoardStore.GetByBoardName(newPost.Title)
				if err == nil {
					// We found a board with this title
					canProceed = false
				} else {
					//Validate the Post
					err = newPost.Validate()
					if err != nil {
						errorMessage = "Error: Could not validate new post"
						status = http.StatusConflict
						canProceed = false
					}
				}
			}
			//don't need to check for duplicate posts as with users
			if canProceed {
				thisPost, err := cr.PostStore.Insert(newPost)
				if err != nil {
					fmt.Printf("Could not insert the post: %v", err)
					status = http.StatusInternalServerError
				}

				err = cr.PostTrie.Insert(thisPost.Title, thisPost.ID, 0)

				if err != nil {
					fmt.Printf("Could not insert the post for search: %v", err)
					status = http.StatusInternalServerError
				} else {
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(&thisPost)
				}
			}
		}
		if !canProceed {
			fmt.Printf(errorMessage + "\n")
			w.WriteHeader(status)
			w.Write([]byte(errorMessage))
		}
	}
}

//UpdatePostHandler should update a post with a PostUpdate
//PATCH /posts/update?id=<id>
func (cr *ContextReceiver) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	//post id delivered as part of post/put/patch request?
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		//status := http.StatusAccepted
		canProceed := true
		//get post by id?
		updates := &posts.PostUpdate{}
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			canProceed = false
		}

		id := r.URL.Query().Get("id")
		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
			canProceed = false
		}
		if canProceed {
			postid := bson.ObjectIdHex(id)

			//get post from db
			post, err := cr.PostStore.GetByID(postid)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				canProceed = false
			}
			//authenticate
			if canProceed {
				sid, err := sessions.GetSessionID(r, cr.SigningKey)
				if err != nil {
					// No we don't have a session so we can't do this
					w.WriteHeader(http.StatusForbidden)
					canProceed = false
				}
				if canProceed {
					state := &SessionState{}
					err := cr.SessionStore.Get(sid, state)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						canProceed = false
					}
					if state.User.ID != post.AuthorID {
						w.WriteHeader(http.StatusForbidden)
						canProceed = false
					}
				}
			}

			//postnew := &posts.Post{post}
			if canProceed {
				err := json.NewDecoder(r.Body).Decode(updates)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					canProceed = false
				}
				//apply updates
				if canProceed {
					//to the struct
					err = post.ApplyUpdates(updates)
					if err != nil {
						fmt.Printf("Could not update post state\n")
						canProceed = false
					}
					//to the user store
					err = cr.PostStore.PostUpdate(postid, updates)
					if err != nil {
						fmt.Printf("Could not update post in database\n")
						canProceed = false
					}
				}
				if canProceed {
					w.WriteHeader(http.StatusAccepted)
					// Output the updated user back
					json.NewEncoder(w).Encode(post)
				}

			}
		}

		if !canProceed {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

//GetPostHandler returns the content of a single post
func (cr *ContextReceiver) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			postid := bson.ObjectIdHex(id)
			post, err := cr.PostStore.GetByID(postid)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(post)

		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GetPostByBoardHandler returns all the posts
func (cr *ContextReceiver) GetPostByBoardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			boardid := bson.ObjectIdHex(id)
			post, err := cr.PostStore.GetByBoardID(boardid)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(post)

		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GetPostByAuthorHandler gets all the posts by the given author
func (cr *ContextReceiver) GetPostByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			authorid := bson.ObjectIdHex(id)
			posts, err := cr.PostStore.GetByAuthorID(authorid)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(posts)

		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GetLastNHandler is a handler for getting the last N posts
func (cr *ContextReceiver) GetLastNHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		n := r.URL.Query().Get("n")
		if len(n) == 0 {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			n, err := strconv.Atoi(n)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				posts, err := cr.PostStore.GetLastN(n)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}
				w.WriteHeader(http.StatusAccepted)
				json.NewEncoder(w).Encode(posts)
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
