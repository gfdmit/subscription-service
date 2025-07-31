package v1

import (
	"net/http"
	"time"

	"github.com/gfdmit/subscription-service/internal/handlers/http/v1/rest"
	"github.com/gfdmit/subscription-service/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(svc service.Service) *gin.Engine {
	var (
		router = gin.New()
	)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300 * time.Second,
	}))

	restHandler := rest.New(svc)

	router.Any("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	api := router.Group("/api")
	{
		api.Use(gin.Logger())
		v1 := api.Group("/v1")
		{
			subscriptions := v1.Group("/subscriptions")
			{
				subscriptions.POST("", restHandler.CreateSubscription)
				subscriptions.GET("/:id", restHandler.GetSubscription)
				subscriptions.GET("", restHandler.GetSubscriptions)
				subscriptions.DELETE("/:id", restHandler.DeleteSubscription)
				subscriptions.PUT("/:id", restHandler.UpdateSubscription)
			}
		}
	}

	return router
}
