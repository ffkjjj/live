package huya

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/live/handler"
)

func GetStream(c *gin.Context) {
	roomId := c.Param("room")
	streamBit := c.Param("streamBit")
	ss := handler.GetStreamData(roomId, streamBit)
	c.JSON(http.StatusOK, gin.H{
		"streamUrls": ss,
	})
}

func GetHuyaLiveCategory(c *gin.Context) {
	action := c.Query("action")
	categoryP := handler.GetHuyaLiveCategory(action)

	c.JSON(http.StatusOK, categoryP)
}
