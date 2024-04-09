package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FecthedResult struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func SearchStoriesAndComment(c *fiber.Ctx) error {
	// get keyword from body
	keyword := c.Query("keyword")
	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid keyword")
	}
	print(keyword)
	stories, err := SearchStoriesByKeyword(keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to search for stories")
	}
	return c.JSON(stories)
}

func SearchStoriesByKeyword(keyword string) ([]FecthedResult, error) {
	// Make a request to the Hacker News API to search for stories
	url := fmt.Sprintf("https://hn.algolia.com/api/v1/search?query=%s&tags=story", keyword)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response into a SearchResponse struct
	var searchResp struct {
		Hits []FecthedResult `json:"hits"`
	}
	err = json.NewDecoder(resp.Body).Decode(&searchResp)
	if err != nil {
		return nil, err
	}

	return searchResp.Hits, nil
}
