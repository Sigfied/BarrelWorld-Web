package one

import (
	"BarrelWorld-Web/app/api/exception"
	"BarrelWorld-Web/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
	path2 "path"
	"strings"
)

// GetOne retrieves an object from the specified bucket and object name.
func GetOne(context *gin.Context, bucketName string, objectName string) {
	// Retrieve the MinIO client
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
		return
	}
	// Get the object from the specified bucket and object name
	object, err := client.GetObject(context, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		// Return an error response if object retrieval fails
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Retrieve object information, including the key and content type
	info, err := object.Stat()
	// Construct the local file path to save the retrieved object
	k := strings.ReplaceAll(info.Key, "/", "\\")
	k = path2.Ext(k)
	path := fmt.Sprintf("%s%s%s", config.FileSavePath, info.ETag, k)
	//config.Log.Info("path:%v", path)
	localFile, err := os.Create(path)
	if _, err = io.Copy(localFile, object); err != nil {
		// Print the error if copying the object to the local file fails
		config.Log.Info("copy error :%v", err.Error())
		return
	}

	// Send the file as a response to the client

	context.File(path)

	// Close the MinIO object and local file when done
	defer func(object *minio.Object) {
		err := object.Close()
		if err != nil {
			// Handle error if closing the MinIO object fails
		}
	}(object)
	defer func(localFile *os.File) {
		err := localFile.Close()
		if err != nil {
			// Handle error if closing the local file fails
			_ = os.Remove(path)
		}
	}(localFile)
}

// PutOne uploads an object to the specified bucket and object name.
func PutOne(context *gin.Context, bucketName string, objectName string, filePath string) {
	// Retrieve the MinIO client
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
		return
	}

	// Open the file specified by the file path
	file, err := os.Open(filePath)
	// Get the file information, including its size
	fileStat, err := file.Stat()
	if err != nil {
		config.Log.Info("fileStat error:{}", err)
		return
	}

	// Upload the file to the specified bucket and object name
	object, err := client.PutObject(context, bucketName, objectName, file, fileStat.Size(), minio.PutObjectOptions{})
	if err != nil {
		// Return an error response if object uploading fails
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return a success response with the uploaded object information
	context.JSON(200, gin.H{
		"object": object,
	})
	err = os.Remove(filePath)
	if err != nil {
		return
	}
}

// StatOne retrieves object information of the specified bucket and object name.
func StatOne(context *gin.Context, bucketName string, objectName string) {
	client, err := config.Minio()
	if err != nil {
		exception.MinClientError(context, err)
		return
	}

	info, err := client.StatObject(context, bucketName, objectName, minio.StatObjectOptions{})

	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"info": info,
	})
}
