package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmark struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string             `json:"username" bson:"username"`
	CommentIDs []string           `json:"commentIds" bson:"commentIds"`
	StoryIDs   []string           `json:"storyIds" bson:"storyIds"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
