package views

import (
	"insight/pkg/application"
	"insight/pkg/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type DeviceViews struct {
	apps *application.Application
}

func NewDeviceViews(apps *application.Application) *DeviceViews {
	views := DeviceViews{}

	views.apps = apps

	return &views
}

func (view *DeviceViews) Create(ctx *gin.Context) {
	response := Response{}

	model := entity.Device{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := view.apps.DeviceApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Message = "Device Created"
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceViews) Get(ctx *gin.Context) {
	response := Response{}

	id := ctx.Param("id")
	device, err := view.apps.DeviceApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Found"
	response.Payload = device

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceViews) List(ctx *gin.Context) {
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

	devices, err := view.apps.DeviceApp.List(ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Found"
	response.Payload = devices

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceViews) Update(ctx *gin.Context) {
	response := Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.Device{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := view.apps.DeviceApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Device Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
