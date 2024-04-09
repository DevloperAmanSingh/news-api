package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DevloperAmanSingh/news-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromApi(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return c.Status(http.StatusBadRequest).SendString("Invalid user ID")
	}
	user, err := GetUserById(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching user")
	}
	return c.JSON(user)
}

func GetUserById(userId string) (*models.ApiUser, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/user/%s.json", userId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user models.ApiUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
