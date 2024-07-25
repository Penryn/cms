package router

import (
	"cms/internal/handler/user"
	c "cms/internal/handler/contact"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){ 
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/login", user.Login)
		api.POST("/register", user.Reg)

		contact := api.Group("/contact")
		{
			contact.GET("", c.GetContact)
			contact.POST("", c.CreateContact)
			contact.PUT("", c.UpdateContact)
			contact.DELETE("", c.DeleteContact)
		}
	}
}