package routers

import (
	"github.com/bitebait/cupcakestore/controllers"
	"github.com/bitebait/cupcakestore/database"
	"github.com/bitebait/cupcakestore/middlewares"
	"github.com/bitebait/cupcakestore/repositories"
	"github.com/bitebait/cupcakestore/services"
	"github.com/bitebait/cupcakestore/session"
	"github.com/bitebait/cupcakestore/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type UserRouter struct {
	userController   controllers.UserController
	userService      services.UserService
	profileService   services.ProfileService
	templateRenderer views.TemplateRenderer
}

func NewUserRouter() *UserRouter {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	templateRenderer := views.NewTemplateRenderer()
	profileRepository := repositories.NewProfileRepository(database.DB)
	profileService := services.NewProfileService(profileRepository)
	userController := controllers.NewUserController(userService, profileService, templateRenderer)

	return &UserRouter{
		userController:   userController,
		userService:      userService,
		profileService:   profileService,
		templateRenderer: templateRenderer,
	}
}

func (r *UserRouter) InstallRouters(app *fiber.App) {
	user := app.Group("/users", cors.New())
	user.Use(middlewares.LoginAndStaffRequired(session.Store, r.userService))
	user.Get("/create", r.userController.RenderCreate)
	user.Post("/create", r.userController.HandlerCreate)
	user.Get("/", r.userController.RenderUsers)
	user.Get("/:id", r.userController.RenderUser)
	user.Post("/update/:id", r.userController.HandlerUpdate)
	user.Get("/delete/:id", r.userController.RenderDelete)
	user.Post("/delete/:id", r.userController.HandlerDelete)
}
