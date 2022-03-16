package main

import (
	"github.com/gin-gonic/gin"
)
func main()  {
	r:=gin.Default()
	r.GET("/v1/topic", func(c *gin.Context) {
		if c.Query("username")==""{
			c.String(300,"获取帖子列表")
		}else {
			c.String(200,"获取用户=%s的帖子列表",c.Param("username"))
		}
	})
	r.GET("/v1/topics/:topic_id", func(c *gin.Context) {
		c.String(200,"获取topicid=%s的帖子",c.Param("topic_id"))
	})

	r.Run(":8080")
}
