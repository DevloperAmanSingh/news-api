package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/DevloperAmanSingh/news-api/internal/database"
	"github.com/DevloperAmanSingh/news-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var result struct {
	StoryIds   []string `bson:"storyIds"`
	CommentIds []string `bson:"commentIds"`
}

func GetBookmarks(c *fiber.Ctx) error {
	// Extract the username from the request params
	username := c.Params("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid username")
	}

	itemType := c.Params("itemType")
	if itemType != "stories" && itemType != "comments" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid item type")
	}
	collection := db.GetBookmarksCollection()
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON([]string{})
	}
	if itemType == "stories" {
		// Fetch details of each story using the Hacker News API
		var stories []models.Story
		for _, storyID := range result.StoryIds {
			id, err := strconv.Atoi(storyID)
			if err != nil {
				// Handle error converting storyID to int (skip the story)
				continue
			}
			story, err := GetStoryDetails(id)
			if err != nil {
				// Handle error fetching story details (skip the story)
				continue
			}
			stories = append(stories, *story)
		}
		return c.JSON(stories)
	} else if itemType == "comments" {
		// Fetch details of each comment using the Hacker News API
		var comments []models.Comment
		for _, commentID := range result.CommentIds {
			id, err := strconv.Atoi(commentID)
			if err != nil {
				// Handle error converting commentID to int (skip the comment)
				continue
			}
			comment, err := GetCommentByID(id)
			if err != nil {
				// Handle error fetching comment details (skip the comment)
				continue
			}
			comments = append(comments, *comment)
		}
		return c.JSON(comments)
	}

	return c.Status(fiber.StatusBadRequest).SendString("Invalid item type")
}

func GetStoryDetails(id int) (*models.Story, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var story models.Story
	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {
		return nil, err
	}

	return &story, nil
}

func GetCommentByID(id int) (*models.Comment, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var comment models.Comment
	err = json.NewDecoder(resp.Body).Decode(&comment)
	if err != nil {
		return nil, err
	}

	// Recursively fetch replies if the comment has kids
	if len(comment.Kids) > 0 {
		replies := make([]models.Comment, 0)
		for _, kidID := range comment.Kids {
			reply, err := GetCommentByID(kidID)
			if err != nil {
				return nil, err
			}
			replies = append(replies, *reply)
		}
		comment.Replies = replies
	}

	return &comment, nil
}
