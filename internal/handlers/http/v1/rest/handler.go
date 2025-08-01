package rest

import (
	"net/http"
	"strconv"

	"github.com/gfdmit/subscription-service/internal/model"
	"github.com/gfdmit/subscription-service/internal/service"
	"github.com/gfdmit/subscription-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type restHandler struct {
	svc service.Service
}

func New(svc service.Service) *restHandler {
	return &restHandler{svc: svc}
}

//	    CreateSubscriptions godoc
//
//		@Summary		Create subscriptions
//		@Description	create subscription
//		@Accept			json
//		@Produce		json
//		@Param			Subscriptions   body	model.Subscription	true	"Create Subscription"
//		@Success		200	{object}	model.Subscription
//		@Failure		400	{object}	error
//		@Failure		404	{object}	error
//		@Failure		500	{object}	error
//		@Router			/subscriptions/{id} [post]
func (rh restHandler) CreateSubscription(c *gin.Context) {
	var subscription model.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		utils.Error(c, err.Error())
		return
	}

	createdSubscription, err := rh.svc.CreateSubscription(c.Request.Context(), subscription)
	if err != nil {
		utils.Error(c, err.Error())
	}

	c.JSON(http.StatusCreated, createdSubscription)
}

func (rh restHandler) GetSubscription(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, "invalid ID")
		return
	}

	subscription, err := rh.svc.GetSubscription(c.Request.Context(), id)
	if err != nil {
		utils.Error(c, err.Error())
	}

	c.JSON(http.StatusOK, subscription)
}

func (rh restHandler) GetSubscriptions(c *gin.Context) {
	subscriptions, err := rh.svc.GetSubscriptions(c.Request.Context())
	if err != nil {
		utils.Error(c, err.Error())
	}

	c.JSON(http.StatusOK, subscriptions)
}

func (rh restHandler) UpdateSubscription(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, "invalid ID")
		return
	}

	var subscription model.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		utils.Error(c, err.Error())
		return
	}

	updatedSubscriptions, err := rh.svc.UpdateSubscription(c.Request.Context(), id, subscription)
	if err != nil {
		utils.Error(c, err.Error())
	}
	c.JSON(http.StatusOK, updatedSubscriptions)
}

func (rh restHandler) DeleteSubscription(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, "invalid ID")
		return
	}

	isDeleted, err := rh.svc.DeleteSubscription(c.Request.Context(), id)
	if err != nil {
		utils.Error(c, err.Error())
	}

	c.JSON(http.StatusOK, isDeleted)
}

func (rh restHandler) GetAmount(c *gin.Context) {
	activeParams := map[string]string{}

	startDate, ok := c.GetQuery("start_date")
	if !ok {
		utils.Error(c, "there no start date")
		return
	} else {
		activeParams["start_date"] = startDate
	}

	endDate, ok := c.GetQuery("end_date")
	if !ok {
		utils.Error(c, "there no end date")
		return
	} else {
		activeParams["end_date"] = endDate
	}

	userID, ok := c.GetQuery("user_id")
	if ok {
		activeParams["user_id"] = userID
	}

	service, ok := c.GetQuery("service_name")
	if ok {
		activeParams["service"] = service
	}

	amount, err := rh.svc.GetAmount(c.Request.Context(), activeParams)
	if err != nil {
		utils.Error(c, err.Error())
	}

	c.JSON(http.StatusOK, amount)
}
