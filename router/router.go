package router

import (
	"github.com/gin-gonic/gin"
	"web/live/apis"
	"web/live/apis/huya"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", apis.IndexApi)
	{
		liveGroup := router.Group("/live")
		{
			hyLive := liveGroup.Group("/hy")
			hyLive.GET("/list", huya.GetHuyaLiveCategory)
			hyLive.GET("/stream/:room", huya.GetStream)
			hyLive.GET("/stream/:room/:streamBit", huya.GetStream)
		}
	}
	return router
}
