package routers

import (
	"github.com/RicardoSouza26238896/cupcakestore/controllers"
	"github.com/RicardoSouza26238896/cupcakestore/database"
	"github.com/RicardoSouza26238896/cupcakestore/middlewares"
	"github.com/RicardoSouza26238896/cupcakestore/repositories"
	"github.com/RicardoSouza26238896/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	authController controllers.AuthController
}

func NewAuthRouter() *AuthRouter {
	// Initialize repositories
	userRepository := repositories.NewUserRepository(database.DB)
	profileRepository := repositories.NewProfileRepository(database.DB)

	// Initialize services with repositories
	userService := services.NewUserService(userRepository)
	profileService := services.NewProfileService(profileRepository)
	authService := services.NewAuthService(userService, profileService)

	// Initialize controllers with services
	authController := controllers.NewAuthController(authService)

	return &AuthRouter{
		authController: authController,
	}
}

func (r *AuthRouter) InstallRouters(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Get("/login", r.authController.RenderLogin)
	auth.Post("/login", r.authController.Login)
	auth.Get("/register", r.authController.RenderRegister)
	auth.Post("/register", r.authController.Register)
	auth.Get("/logout", r.authController.Logout).Use(middlewares.LoginRequired())
}
