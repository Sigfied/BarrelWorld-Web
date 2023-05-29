package router

import (
	"BarrelWorld-Web/app/api/bucket"
	"BarrelWorld-Web/app/api/object"
	"BarrelWorld-Web/app/api/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	{
		//Get请求路由
		{
			userGroup.GET("/login", user.Login)
		}
	}

	bucketGroup := engine.Group("/bucket")
	{
		//Get请求路由
		{
			bucketGroup.GET("/list", bucket.ListBuckets)
			bucketGroup.GET("/create", bucket.CreateBucket)
			bucketGroup.GET("/all", bucket.AllElement)
		}
	}

	objectGroup := engine.Group("/object")
	{
		//Get请求路由
		{
			objectGroup.GET("/list", object.GetList)
			objectGroup.GET("/one", object.GetOne)
			objectGroup.GET("/stat", object.StatOne)
		}
		//Post请求路由
		{
			objectGroup.POST("/one", object.PutOne)
		}
	}
}
