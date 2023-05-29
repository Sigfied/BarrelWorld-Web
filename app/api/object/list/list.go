package list

import (
	"BarrelWorld-Web/app/api/exception"
	"BarrelWorld-Web/config"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Objects(context *gin.Context, name string, prefix string) {
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
	}
	objectInfos := client.ListObjects(context, name, minio.ListObjectsOptions{
		WithMetadata: true,
		Prefix:       prefix,
		Recursive:    false,
	})

	var data []minio.ObjectInfo
	for obj := range objectInfos {
		data = append(data, obj)
	}

	context.JSON(200, gin.H{
		"objects": data,
		"error":   err,
	})
}
