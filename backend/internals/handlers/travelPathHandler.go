package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/alarm-tool/internals/services"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
	"github.com/iacopoghilardi/alarm-tool/utils"
)

type TravelPathHandler struct {
	travelPathService *services.TravelPathService
}

func NewTravelPathHandler(travelPathService *services.TravelPathService) *TravelPathHandler {
	return &TravelPathHandler{travelPathService: travelPathService}
}

func (h *TravelPathHandler) FindAll(c *gin.Context) {
	travelPaths, err := h.travelPathService.FindAll(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(travelPaths))
}

func (h *TravelPathHandler) Create(c *gin.Context) {
	var travelPathDto dto.CreateTravelPathDto
	if err := c.ShouldBindJSON(&travelPathDto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	travelPath, err := h.travelPathService.Create(&travelPathDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusCreated, utils.BuildSuccessResponse(travelPath))
}

func (h *TravelPathHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var travelPathDto dto.CreateTravelPathDto
	if err := c.ShouldBindJSON(&travelPathDto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}
	travelPath, err := h.travelPathService.Update(id, &travelPathDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(travelPath))
}

func (h *TravelPathHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.travelPathService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
	}
}
