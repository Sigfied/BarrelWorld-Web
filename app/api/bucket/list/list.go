package list

import (
	"BarrelWorld-Web/app/api/exception"
	"BarrelWorld-Web/app/util"
	"BarrelWorld-Web/config"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"path"

	list "github.com/duke-git/lancet/v2/datastructure/list"
)

type BucketInfo struct {
	Name       string
	Size       int64
	SizeStr    string
	CreateTime string
}

func Buckets(context *gin.Context) {
	minioClient, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
		return
	}
	buckets, err := minioClient.ListBuckets(context)
	if err != nil {
		config.Log.Info("List Buckets Error :%v", err)
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.Log.Info("List Buckets :%v\n", buckets)
	context.JSON(200, gin.H{
		"buckets": buckets,
	})
}

func BucketsAllElement(context *gin.Context) {
	minioClient, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
		return
	}

	buckets, err := minioClient.ListBuckets(context)

	if err != nil {
		config.Log.Info("List Buckets Error :%v", err)
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var img []minio.ObjectInfo
	var doc []minio.ObjectInfo
	var oth []minio.ObjectInfo
	var vid []minio.ObjectInfo

	bucketsList := list.NewList(buckets)

	bucketsInfoArray := make([]BucketInfo, len(buckets))

	sum := 0
	bucketsList.ForEach(func(info minio.BucketInfo) {
		objects := minioClient.ListObjects(context, info.Name, minio.ListObjectsOptions{
			Recursive: true,
		})
		if err != nil {
			config.Log.Info("List Buckets Error :%v", err)
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		size := int64(0)

		for object := range objects {
			size += object.Size
			t := path.Ext(object.Key)
			if config.TypeMap.GetFileType(t) != nil {
				if config.TypeMap.GetFileType(t) == config.ImageType {
					img = append(img, object)
				} else if config.TypeMap.GetFileType(t) == config.VideoType {
					vid = append(vid, object)
				} else if config.TypeMap.GetFileType(t) == config.DocumentType {
					doc = append(doc, object)
				} else if config.TypeMap.GetFileType(t) == config.OtherType {
					oth = append(oth, object)
				}
			}
			sum++
		}

		bucketsInfoArray = append(bucketsInfoArray, BucketInfo{
			Name:       info.Name,
			Size:       size,
			SizeStr:    util.FormatFileSize(size),
			CreateTime: info.CreationDate.String(),
		})
	})

	for i := 0; i < len(bucketsInfoArray); i++ {
		if bucketsInfoArray[i].Name == "" {
			bucketsInfoArray = append(bucketsInfoArray[:i], bucketsInfoArray[i+1:]...)
			i--
		}
	}

	context.JSON(200, gin.H{
		"buckets":   bucketsInfoArray,
		"all":       sum,
		"images":    img,
		"imagesNum": len(img),
		"videos":    vid,
		"videosNum": len(vid),
		"documents": doc,
		"docNum":    len(doc),
		"others":    oth,
		"othNum":    len(oth),
	})
}
