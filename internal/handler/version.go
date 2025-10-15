package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/versegeek/go-skeleton/pkg/version"
)

type VersionHandler interface {
	VersionEndpoint(ginCtx *gin.Context)
}

func (h *handler) VersionEndpoint(ginCtx *gin.Context) {
	ginCtx.AbortWithStatusJSON(http.StatusOK, version.Get().ToString())
	return
}
