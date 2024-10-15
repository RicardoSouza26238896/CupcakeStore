package routers

import (
	"github.com/RicardoSouza26238896/cupcakestore/controllers"
	"github.com/RicardoSouza26238896/cupcakestore/database"
	"github.com/RicardoSouza26238896/cupcakestore/middlewares"
	"github.com/RicardoSouza26238896/cupcakestore/repositories"
	"github.com/RicardoSouza26238896/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type ShoppingCartRouter struct {
	shoppingCartController controllers.ShoppingCartController
}

func NewShoppingCartRouter() *ShoppingCartRouter {
	// Initialize repositories
	shoppingCartRepository := repositories.NewShoppingCartRepository(database.DB)
	shoppingCartItemRepository := repositories.NewShoppingCartItemRepository(database.DB)

	// Initialize services with repositories
	shoppingCartItemService := services.NewShoppingCartItemService(shoppingCartItemRepository)
	shoppingCartService := services.NewShoppingCartService(shoppingCartRepository, shoppingCartItemService)

	// Initialize controllers with services
	shoppingCartController := controllers.NewShoppingCartController(shoppingCartService)

	return &ShoppingCartRouter{
		shoppingCartController: shoppingCartController,
	}
}

func (r *ShoppingCartRouter) InstallRouters(app *fiber.App) {
	cart := app.Group("/cart").Use(middlewares.LoginRequired())
	cart.Get("/", r.shoppingCartController.RenderShoppingCart)
	cart.Post("/", r.shoppingCartController.AddShoppingCartItem)
	cart.Get("/count", r.shoppingCartController.CountShoppingCart)
	cart.Get("/remove/:id", r.shoppingCartController.RemoveFromCart)
}
