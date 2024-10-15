package routers

import (
	"github.com/RicardoSouza26238896/cupcakestore/controllers"
	"github.com/RicardoSouza26238896/cupcakestore/database"
	"github.com/RicardoSouza26238896/cupcakestore/middlewares"
	"github.com/RicardoSouza26238896/cupcakestore/repositories"
	"github.com/RicardoSouza26238896/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type OrderRouter struct {
	orderController controllers.OrderController
}

func NewOrderRouter() *OrderRouter {
	// Initialize repositories
	orderRepository := repositories.NewOrderRepository(database.DB)
	storeConfigRepository := repositories.NewStoreConfigRepository(database.DB)

	// Initialize services with repositories
	storeConfigService := services.NewStoreConfigService(storeConfigRepository)
	orderService := services.NewOrderService(orderRepository, storeConfigService)

	// Initialize controllers with services
	orderController := controllers.NewOrderController(orderService, storeConfigService)

	return &OrderRouter{
		orderController: orderController,
	}
}

func (r *OrderRouter) InstallRouters(app *fiber.App) {
	order := app.Group("/orders").Use(middlewares.LoginRequired())
	order.Get("/checkout/:id", r.orderController.Checkout)
	order.Post("/payment/:id", r.orderController.Payment)
	order.Get("/payment/:id", r.orderController.Payment)
	order.Get("/cancel/:id", r.orderController.RenderCancel)
	order.Post("/cancel/:id", r.orderController.Cancel)
	order.Get("/", r.orderController.RenderAllOrders)
	order.Get("/order/:id", r.orderController.RenderOrder)
	order.Post("/order/:id", r.orderController.Update)
}
