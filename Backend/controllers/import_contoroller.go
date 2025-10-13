package controllers

import (
	"main/middleware"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IImportController interface {
	ImportCSV(ctx *gin.Context)
}

type ImportController struct {
	service services.IImportService
}

func NewImportController(service services.IImportService) IImportController {
	return &ImportController{service: service}
}

func (c *ImportController) ImportCSV(ctx *gin.Context) {
	userId := ctx.GetUint("sub")

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		middleware.Log.Error("failed to get file", zap.Error(err))
		return
	}

	if err := c.service.ImportCSV(userId, file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		middleware.Log.Error("failed to import csv", zap.Error(err))
		return
	}
	ctx.Status(http.StatusOK)
}
