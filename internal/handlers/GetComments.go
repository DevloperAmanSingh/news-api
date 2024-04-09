package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DevloperAmanSingh/news-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

var story models.Story

func GetStoryComments(c *fiber.Ctx) error {
	// Get the story ID from the URL parameter
	storyID := c.Params("id")

	// Construct the URL for fetching story details
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", storyID)

	// Fetch story details from Hacker News API
	resp, err := http.Get(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching story details")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&story)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error decoding story details")
	}

	// Check if the story is a story type
	if story.Type != "story" {
		return c.Status(http.StatusBadRequest).SendString("Invalid story ID")
	}

	// Fetch comment IDs from the story
	commentIDs := story.Kids
	comments := make([]models.Comment, 0)

	// Fetch each comment by ID
	for _, id := range commentIDs {
		comment, err := GetCommentByIDs(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error fetching comments")
		}
		comments = append(comments, *comment)
	}

	// Return comments as JSON response
	return c.JSON(comments)
}

func GetCommentByIDs(id int) (*models.Comment, error) {
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
			reply, err := GetCommentByIDs(kidID)
			if err != nil {
				return nil, err
			}
			replies = append(replies, *reply)
		}
		comment.Replies = replies
	}

	return &comment, nil
}
