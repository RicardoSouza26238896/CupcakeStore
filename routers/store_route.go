package routers

import (
	"github.com/RicardoSouza26238896/cupcakestore/controllers"
	"github.com/RicardoSouza26238896/cupcakestore/database"
	"github.com/RicardoSouza26238896/cupcakestore/repositories"
	"github.com/RicardoSouza26238896/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type StoreRouter struct {
	storeController controllers.StoreController
}

func NewStoreRouter() *StoreRouter {
	// Initialize repositories
	productRepository := repositories.NewProductRepository(database.DB)

	// Initialize services with repositories
	productService := services.NewProductService(productRepository)

	// Initialize controllers with services
	storeController := controllers.NewStoreController(productService)
	return &StoreRouter{
		storeController: storeController,
	}
}

func (r *StoreRouter) InstallRouters(app *fiber.App) {
	store := app.Group("/store")
	store.Get("/", r.storeController.RenderStore)
}
