package views

import (
	"insight/pkg/application"
	"insight/pkg/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type DeviceDataViews struct {
	apps *application.Application
}

func NewDeviceDataViews(apps *application.Application) *DeviceDataViews {
	views := DeviceDataViews{}

	views.apps = apps

	return &views
}

func (view *DeviceDataViews) Create(ctx *gin.Context) {
	response := Response{}

	model := entity.DeviceData{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := view.apps.DeviceDataApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Created"
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceDataViews) Get(ctx *gin.Context) {
	response := Response{}

	id := ctx.Param("id")
	deviceData, err := view.apps.DeviceDataApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Found"
	response.Payload = deviceData

	ctx.JSON(http.StatusOK, response)
}

func (view *DeviceDataViews) List(ctx *gin.Context) {
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

	deviceData, err := view.apps.DeviceDataApp.List(ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Found"
	response.Payload = deviceData

	ctx.JSON(http.StatusOK, response)
}
