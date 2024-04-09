package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/DevloperAmanSingh/news-api/internal/database"
	"github.com/DevloperAmanSingh/news-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tempBookmark struct {
	Username string `json:"username" bson:"username"`
	ItemID   string `json:"itemId" bson:"itemId"`
}

func SetBookmarks(c *fiber.Ctx) error {

	var bookmark tempBookmark
	if err := c.BodyParser(&bookmark); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	postType := GetPostTypeFromID(bookmark.ItemID)
	if postType == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid post ID")
	}

	bookMarkType := GetPostTypeFromID(bookmark.ItemID)
	if bookMarkType == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid post ID")
	}
	collection := db.GetBookmarksCollection()
	if bookMarkType == "comment" {
		filter := bson.M{"username": bookmark.Username}
		update := bson.M{"$addToSet": bson.M{"commentIds": bookmark.ItemID}}
		collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	} else if bookMarkType == "story" {
		filter := bson.M{"username": bookmark.Username}
		update := bson.M{"$addToSet": bson.M{"storyIds": bookmark.ItemID}}
		collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	}

	// Save the bookmark to the database
	return c.JSON(bookmark)
}

func GetPostTypeFromID(id string) string {
	var Url = "https://hacker-news.firebaseio.com/v0/item/" + id + ".json"
	resp, err := http.Get(Url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	var postType models.Story
	if err := json.NewDecoder(resp.Body).Decode(&postType); err != nil {
		return ""
	}
	fmt.Print(postType.Type)
	return postType.Type
}
