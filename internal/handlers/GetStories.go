package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DevloperAmanSingh/news-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

var Story models.Story

func GetStories(c *fiber.Ctx) error {
	// Get the category query parameter
	category := c.Query("category")

	// Get the pagination parameters
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Determine the URL based on the category
	url := "https://hacker-news.firebaseio.com/v0/"
	switch category {
	case "top":
		url += "topstories.json"
	case "new":
		url += "newstories.json"
	case "best":
		url += "beststories.json"
	default:
		url += "topstories.json" // Default to top stories if no or invalid category provided
	}

	// Fetch story IDs from Hacker News API
	storyIDs, err := GetStoryIDs(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching story IDs")
	}

	// Fetch stories based on IDs with pagination
	stories, err := GetStoriesByID(storyIDs, pageInt, limitInt)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching stories")
	}

	// Return stories as JSON response
	return c.JSON(stories)
}

func GetStoryIDs(url string) ([]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storyIDs []int
	err = json.NewDecoder(resp.Body).Decode(&storyIDs)
	if err != nil {
		return nil, err
	}

	return storyIDs, nil
}

func GetStoriesByID(ids []int, page, limit int) ([]models.Story, error) {
	var stories []models.Story
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit
	if endIndex > len(ids) {
		endIndex = len(ids)
	}
	for _, id := range ids[startIndex:endIndex] {
		story, err := GetStoryByID(id)
		if err != nil {
			return nil, err
		}
		stories = append(stories, *story)
	}
	return stories, nil
}

func GetStoryByID(id int) (*models.Story, error) {
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
