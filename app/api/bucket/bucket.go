package bucket

import (
	"BarrelWorld-Web/app/api/bucket/get"
	"BarrelWorld-Web/app/api/bucket/list"
	"github.com/gin-gonic/gin"
)

func ListBuckets(context *gin.Context) {
	list.Buckets(context)
}

func CreateBucket(ctx *gin.Context) {
	bucketName, flag := ctx.GetQuery("name")
	if !flag {
		ctx.JSON(400, gin.H{
			"error": "name is required",
			"flag":  false,
		})
	}
	get.Create(ctx, bucketName)
}

func AllElement(ctx *gin.Context) {
	list.BucketsAllElement(ctx)
}
