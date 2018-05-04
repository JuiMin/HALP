package handlers

import (
	"fmt"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"
	"github.com/JuiMin/HALP/servers/gateway/models/comments"
	"github.com/JuiMin/HALP/servers/gateway/models/posts"
	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

// ContextReceiver keeps track of the Various Storage services and keeps a reference to each of them
type ContextReceiver struct {
	SigningKey   string
	UserStore    users.Store
	SessionStore sessions.Store
	PostStore    posts.Store
	BoardStore   boards.Store
	CommentStore comments.Store
}

// NewContextReceiver Creates a new context receiver with a session key and references to the session and user stores
func NewContextReceiver(key string,
	userStore users.Store,
	redisStore sessions.Store,
	commentStore comments.Store,
	postStore posts.Store,
	boardStore boards.Store) (*ContextReceiver, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("No key set for signing key")
	}
	return &ContextReceiver{
		SigningKey:   key,
		UserStore:    userStore,
		SessionStore: redisStore,
		PostStore:    postStore,
		BoardStore:   boardStore,
		CommentStore: commentStore,
	}, nil
}
