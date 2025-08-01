package v1

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateSubscription(c *gin.Context)
	GetSubscription(c *gin.Context)
	GetSubscriptions(c *gin.Context)
	UpdateSubscription(c *gin.Context)
	DeleteSubscription(c *gin.Context)
	GetAmount(c *gin.Context)
}
