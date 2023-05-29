package object

import (
	"BarrelWorld-Web/app/api/object/list"
	"BarrelWorld-Web/app/api/object/one"
	"BarrelWorld-Web/config"
	"github.com/gin-gonic/gin"
	"log"
)

func GetList(context *gin.Context) {
	name, flag := context.GetQuery("bucketName")
	prefix, _ := context.GetQuery("prefix")
	if !flag {
		context.JSON(400, gin.H{
			"error": "BucketName is required",
		})
	}
	list.Objects(context, name, prefix)
}

func GetOne(context *gin.Context) {
	log.Printf("GetOne")
	bucketName, flagBucket := context.GetQuery("bucketName")
	objectName, flagObject := context.GetQuery("objectName")

	if !flagBucket || !flagObject {
		context.JSON(400, gin.H{
			"error": "bucketName and objectName is required",
		})
		return
	}
	one.GetOne(context, bucketName, objectName)
}

func PutOne(context *gin.Context) {
	bucketName := context.PostForm("bucketName")
	objectName := context.PostForm("objectName")

	file, _ := context.FormFile("file")
	dst := config.FileSavePath + file.Filename
	// 上传文件至指定的完整文件路径
	_ = context.SaveUploadedFile(file, dst)
	if bucketName == "" || objectName == "" {
		context.JSON(400, gin.H{
			"error": "bucketName and objectName is required",
		})
		return
	}

	one.PutOne(context, bucketName, objectName, dst)

}

func StatOne(context *gin.Context) {
	bucketName, flagBucket := context.GetQuery("bucketName")
	objectName, flagObject := context.GetQuery("objectName")

	if !flagBucket || !flagObject {
		context.JSON(400, gin.H{
			"error": "bucketName and objectName is required",
		})
		return
	}
	one.StatOne(context, bucketName, objectName)
}
