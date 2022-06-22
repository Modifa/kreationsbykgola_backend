package controller

import (
	"Modifa/kreationsbykgola_backend/models"
	s "Modifa/kreationsbykgola_backend/services"
	v "Modifa/kreationsbykgola_backend/views"
	"os"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*Email ...*/
func Order(c *gin.Context) {
	var u models.Order

	vr := *new(v.Messages).New()

	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	message := strings.ReplaceAll(vr.OrderMessage, "{FullName}", u.Name)
	message = strings.ReplaceAll(message, "{Product}", u.Product)
	message = strings.ReplaceAll(message, "{Quantity}", u.Quantity)
	message = strings.ReplaceAll(message, "{Flavour}", u.Flavour)
	message = strings.ReplaceAll(message, "{Email}", u.Email)
	message = strings.ReplaceAll(message, "{PhoneNumber}", u.PhoneNumber)
	message = strings.ReplaceAll(message, "{Address}", u.Address)

	//Send Order to Baker Team
	To := os.Getenv("KreationsByKgolaEmail")
	s.OrderRecieved(To, message, u.Name)

	//Send None Response Message To Customer
	s.OrderRecievedResponse(u.Email)

	//Record Order On DB
}
