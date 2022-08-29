package routes

import (
	"Satu/controllers"
	"Satu/middlewares"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddlewares(e)
	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJWT.GET("/customers", controllers.GetCustomersController)
	eJWT.GET("/customers/:id", controllers.GetCustomerByIdController)
	eJWT.POST("/customers", controllers.CreateCustomerController)
	eJWT.PUT("/customers/:id", controllers.UpdateCustomerByIdController)
	eJWT.DELETE("/customers/:id", controllers.DeleteCustomerByIdController)

	e.POST("/login", controllers.LoginCustomerController)

	eJWT.GET("/orders", controllers.GetOrdersController)
	eJWT.GET("/orders/:id", controllers.GetOrderByIdController)
	eJWT.POST("/orders", controllers.CreateOrderController)
	eJWT.PUT("/orders/:id", controllers.UpdateOrderByIdController)
	eJWT.DELETE("/orders/:id", controllers.DeleteOrderByIdController)

	eJWT.GET("/order_details", controllers.GetOrderDetailsController)
	eJWT.GET("/order_details/:id", controllers.GetOrderDetailByIdController)
	eJWT.POST("/order_details", controllers.CreateOrderDetailController)
	eJWT.PUT("/order_details/:id", controllers.UpdateOrderDetailByIdController)
	eJWT.DELETE("/order_details/:id", controllers.DeleteOrderDetailByIdController)

	eJWT.GET("/payment_methods", controllers.GetPaymentMethodsController)
	eJWT.GET("/payment_methods/:id", controllers.GetPaymentMethodByIdController)
	eJWT.POST("/payment_methods", controllers.CreatePaymentMethodController)
	eJWT.PUT("/payment_methods/:id", controllers.UpdatePaymentMethodByIdController)
	eJWT.DELETE("/payment_methods/:id", controllers.DeletePaymentMethodByIdController)

	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:id", controllers.GetProductByIdController)
	eJWT.POST("/products", controllers.CreateProductController)
	eJWT.PUT("/products/:id", controllers.UpdateProductByIdController)
	eJWT.DELETE("/products/:id", controllers.DeleteProductByIdController)

	e.POST("/register", controllers.RegisterCustomerController)

	e.POST("/register_user", controllers.RegisterUserController)
	e.POST("/login_user", controllers.LoginUserController)

	eJWT.GET("/merchants", controllers.GetMerchantsController)
	eJWT.GET("/merchants/:id", controllers.GetMerchantByIdController)
	eJWT.POST("/merchants", controllers.CreateMerchantController)
	eJWT.PUT("/merchants/:id", controllers.UpdateMerchantByIdController)
	eJWT.DELETE("/merchants/:id", controllers.DeleteMerchantByIdController)

	eJWT.GET("/outlets", controllers.GetOutletsController)
	eJWT.GET("/outlets/:id", controllers.GetOutletsController)
	eJWT.POST("/outlets", controllers.CreateOutletController)
	eJWT.PUT("/outlets/:id", controllers.UpdateOutletByIdController)
	eJWT.DELETE("/outlets/:id", controllers.DeleteOutletByIdController)

	eJWT.GET("/transactions", controllers.GetTransactionsController)
	eJWT.GET("/transactions/:id", controllers.GetTransactionByIdController)
	eJWT.POST("/transactions", controllers.CreateTransactionController)
	eJWT.PUT("/transactions/:id", controllers.UpdateTransactionByIdController)
	eJWT.DELETE("/transactions/:id", controllers.DeleteTransactionByIdController)

	return e
}
