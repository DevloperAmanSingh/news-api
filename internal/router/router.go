package router

import (
	"github.com/DevloperAmanSingh/news-api/internal/controllers"
	"github.com/DevloperAmanSingh/news-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	// Define routes
	app.Get("/api/stories", func(c *fiber.Ctx) error {
		return handlers.GetStories(c)
	})

	app.Get("/api/stories/:id/comments", func(c *fiber.Ctx) error {
		return handlers.GetStoryComments(c)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.SignUp(c)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c)
	})

	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		return handlers.GetUserFromApi(c)
	})

	app.Post("/api/bookmarks", func(c *fiber.Ctx) error {
		return handlers.SetBookmarks(c)
	})
	app.Get("/api/bookmarks/:username/:itemType", func(c *fiber.Ctx) error {
		return handlers.GetBookmarks(c)
	})
	app.Get("/api/search", func(c *fiber.Ctx) error {
		return handlers.SearchStoriesAndComment(c)
	})
	return app
}
