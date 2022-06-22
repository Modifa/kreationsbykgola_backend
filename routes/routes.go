package routes

import (
	c "Modifa/kreationsbykgola_backend/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//Pay
	Orders := r.Group("/api/orders/")
	{
		Orders.POST("Order", c.Order)
	}
}
