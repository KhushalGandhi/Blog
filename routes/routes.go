package routes

import (
	"Blog/controllers"
	"Blog/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// User routes
	app.Post("/register", controllers.Register) // done
	app.Post("/login", controllers.Login)       // done
	app.Put("/profile", middlewares.ProtectedRoute, controllers.UpdateProfile)
	app.Get("/profile", middlewares.ProtectedRoute, controllers.ViewProfile) // done
	// Blog post routes
	app.Post("/posts", middlewares.ProtectedRoute, controllers.CreatePost)       // done
	app.Get("/posts", controllers.GetAllPosts)                                   // done
	app.Get("/posts/:id", controllers.GetPost)                                   // done
	app.Put("/posts/:id", middlewares.ProtectedRoute, controllers.UpdatePost)    // done
	app.Delete("/posts/:id", middlewares.ProtectedRoute, controllers.DeletePost) // done
	//app.Get("/posts", controllers.GetAllPosts)

}
