package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/alarm-tool/internals/services"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
	"github.com/iacopoghilardi/alarm-tool/utils"
)

type AlarmHandler struct {
	service *services.AlarmService
}

func NewAlarmHandler(alarmService *services.AlarmService) *AlarmHandler {
	return &AlarmHandler{service: alarmService}
}

func (h *AlarmHandler) FindAll(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	alarms, err := h.service.FindAll(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(alarms))
}

func (h *AlarmHandler) FindById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	alarm, err := h.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(alarm))
}

func (h *AlarmHandler) Create(c *gin.Context) {
	var dto dto.CreateAlarmDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	alarm, err := h.service.Create(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(alarm))
}

func (h *AlarmHandler) Update(c *gin.Context) {
	var dto dto.UpdateAlarmDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}
}

func (h *AlarmHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	err = h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
	}
}
