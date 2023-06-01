package views

import (
	"insight/pkg/application"
	"insight/pkg/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type VesselViews struct {
	apps *application.Application
}

func NewVesselViews(apps *application.Application) *VesselViews {
	views := VesselViews{}

	views.apps = apps

	return &views
}

func (view *VesselViews) Create(ctx *gin.Context) {
	response := Response{}

	model := entity.Vessel{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := view.apps.VesselApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = nil

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Message = "Vessel Created"
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (view *VesselViews) Get(ctx *gin.Context) {
	response := Response{}

	id := ctx.Param("id")
	vessel, err := view.apps.VesselApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Vessel Found"
	response.Payload = vessel

	ctx.JSON(http.StatusOK, response)
}

func (view *VesselViews) List(ctx *gin.Context) {
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

	vessels, err := view.apps.VesselApp.List(ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Vessels Found"
	response.Payload = vessels

	ctx.JSON(http.StatusOK, response)
}

func (view *VesselViews) Update(ctx *gin.Context) {
	response := Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.Vessel{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := view.apps.VesselApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Vessel Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
