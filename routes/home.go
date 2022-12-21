package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to Go Modular")
}
