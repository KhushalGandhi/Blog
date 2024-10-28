package routes

import (
	"Blog/controllers"
	"Blog/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// User routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Put("/profile", middlewares.ProtectedRoute, controllers.UpdateProfile)
	// Blog post routes
	app.Post("/posts", middlewares.ProtectedRoute, controllers.CreatePost)
	app.Get("/posts", controllers.GetPost)
	app.Get("/posts/:id", controllers.GetPost)
	app.Put("/posts/:id", middlewares.ProtectedRoute, controllers.UpdatePost)
	app.Delete("/posts/:id", middlewares.ProtectedRoute, controllers.DeletePost)
	app.Get("/posts", controllers.GetAllPosts)

}
