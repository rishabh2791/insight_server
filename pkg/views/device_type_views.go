package views

import (
	"insight/pkg/application"
	"insight/pkg/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type DeviceTypeViews struct {
	apps *application.Application
}

func NewDeviceTypeViews(apps *application.Application) *DeviceTypeViews {
	views := DeviceTypeViews{}

	views.apps = apps

	return &views
}

func (view *DeviceTypeViews) Create(ctx *gin.Context) {
	response := Response{}

	model := entity.DeviceType{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := view.apps.DeviceTypeApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Message = "Device Type Created"
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceTypeViews) Get(ctx *gin.Context) {
	response := Response{}

	id := ctx.Param("id")
	deviceType, err := view.apps.DeviceTypeApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Type Found"
	response.Payload = deviceType

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceTypeViews) List(ctx *gin.Context) {
	response := Response{}

	conditions := map[string]interface{}{}
	jsonError := json.NewDecoder(ctx.Request.Body).Decode(&conditions)
	if jsonError != nil {
		response.Status = false
		response.Message = jsonError.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	deviceTypes, err := view.apps.DeviceTypeApp.List(ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Types Found"
	response.Payload = deviceTypes

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceTypeViews) Update(ctx *gin.Context) {
	response := Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.DeviceType{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := view.apps.DeviceTypeApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Device Type Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
