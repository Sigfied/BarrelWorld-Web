package exception

import (
	"BarrelWorld-Web/config"
	"github.com/gin-gonic/gin"
)

func MinClientError(context *gin.Context, err error) {
	config.Log.Info("Minio Client Error :%v\n", err)
	context.JSON(500, gin.H{
		"error": err.Error(),
	})
	return
}
