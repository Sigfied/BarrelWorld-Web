package get

import (
	"BarrelWorld-Web/app/api/exception"
	"BarrelWorld-Web/config"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Create(context *gin.Context, name string) {
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
	}
	err = client.MakeBucket(context, name, minio.MakeBucketOptions{Region: "cn-north-1"})
	if err != nil {
		config.Log.Info("Make Bucket Error :%v\n", err)
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	context.JSON(200, gin.H{
		"flag":  true,
		"error": err,
	})

}
