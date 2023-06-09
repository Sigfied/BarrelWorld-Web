package list

import (
	"BarrelWorld-Web/app/api/exception"
	"BarrelWorld-Web/app/util"
	"BarrelWorld-Web/config"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Objects(context *gin.Context, name string, prefix string) {
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
	}

	var data []minio.ObjectInfo
	config.Log.Info("name %v ,prefix %v", name, prefix)
	if prefix != "" {
		objectInfos := client.ListObjects(context, name, minio.ListObjectsOptions{
			WithMetadata: true,
			Prefix:       prefix,
			Recursive:    false,
		})
		for obj := range objectInfos {
			data = append(data, obj)
		}
	} else {
		objectInfos := client.ListObjects(context, name, minio.ListObjectsOptions{
			WithMetadata: true,
			Recursive:    false,
		})
		for obj := range objectInfos {
			data = append(data, obj)
		}
	}
	context.JSON(200, gin.H{
		"objects": data,
		"error":   err,
	})
}

func FuzzySearch(context *gin.Context, name string, sub string) {
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
	}
	objectInfos := client.ListObjects(context, name, minio.ListObjectsOptions{
		WithMetadata: true,
		//Prefix:       prefix,
		Recursive: true,
	})

	var data []minio.ObjectInfo
	for obj := range objectInfos {
		data = append(data, obj)
	}

	search := util.FuzzySearch(sub, data)

	context.JSON(200, gin.H{
		"objects": search,
		"error":   err,
	})
}

func LatestObjects(context *gin.Context, name string, prefix string) {
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
	}
	objectInfos := client.ListObjects(context, name, minio.ListObjectsOptions{
		WithMetadata: true,
		//Prefix:       prefix,
		Recursive: false,
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
